package auth

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/scloud/version"

	cf "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/config"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const (
	defaultScheme  = "https"
	defaultPort    = "443"
	defaultTimeout = 10
	// use ScloudTestDone to indicate the scloud test command is done when using testhook-dryrun flag
	ScloudTestDone = "Test Command is done\n"
	//Switch to turn on and off tenant scoped hostname for auth and api domains
	enableTenantScope bool = false
)

var sdkClient *sdk.Client

func GetClient() (*sdk.Client, error) {
	if sdkClient == nil {
		err := loadConfigs()
		if err != nil {
			return nil, err
		}

		sdkClient = apiClient()

		if sdkClient == nil {
			return nil, errors.New("no valid sdk client")
		}

		return sdkClient, err
	}

	return sdkClient, nil
}

func GetClientSystemTenant() (*sdk.Client, error) {
	if sdkClient == nil {
		err := loadConfigs()
		if err != nil {
			return nil, err
		}

		sdkClient = apiClientWithTenant("system")

		return sdkClient, nil
	}

	return sdkClient, nil
}

func getTimeoutForConfig() time.Duration {
	defaultTimeoutForConfig := time.Duration(defaultTimeout) * time.Second

	timeout, _ := cf.GlobalFlags["timeout"].(uint)
	if timeout != 0 {
		return time.Duration(timeout) * time.Second
	}

	if timeoutInSettings, ok := settings.GetString("timeout"); ok && timeoutInSettings != "" {
		timeout, err := strconv.Atoi(timeoutInSettings)
		if err == nil {
			return time.Duration(timeout) * time.Second
		}
	}

	return defaultTimeoutForConfig
}

// Returns a service client ( points to the new SDK Client) based on the given service config.
func newClient(svc *Service) *sdk.Client {
	// enable multi-region hostnames
	var tenantScoped = enableTenantScope

	var serviceURL *url.URL
	var hostURL = getHostURL()
	serviceURL, err := url.Parse(hostURL)
	var scheme string
	if scheme = serviceURL.Scheme; scheme == "" {
		if scheme = svc.Scheme; scheme == "" {
			scheme = defaultScheme
		}
	}

	var port string
	if port = serviceURL.Port(); port == "" {
		if port = svc.Port; port == "" {
			port = defaultPort
		}
	}

	// hostname obtained from default.yaml contains api. prefix that GoSDK client will append
	host := strings.TrimPrefix(svc.Host, "api.")
	overrideHost := serviceURL.Hostname()

	var overrideHostPort string
	var hostPort string
	if overrideHost != "" {
		overrideHostPort = overrideHost + ":" + port
	} else {
		hostPort = host + ":" + port
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: isInsecure()}

	region := getRegion()
	tenantScopedSetting := getTenantScoped()
	if tenantScopedSetting != false {
		tenantScoped = tenantScopedSetting
	}

	// Load client cert
	caCert := getCaCert()

	// -insecure=false -scheme=https -ca-cert=<path-to-file.crt>
	if !isInsecure() && scheme == defaultScheme && caCert != "" {
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		certs, err := ioutil.ReadFile(caCert)
		if err != nil {
			util.Warning("Failed to append %q to RootCAs: %v", caCert, err)
		}
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			util.Warning("No certs appended, using system certs only")
		}
		// set the RootCA
		tlsConfig.RootCAs = rootCAs
	}

	var scloudVersion string
	scloudVersion = fmt.Sprintf("%s/%s", version.UserAgent, version.ScloudVersion)

	var roundTripper http.RoundTripper

	roundTripper = util.NewCustomSdkTransport(&GlogWrapper{}, &http.Transport{
		TLSClientConfig: tlsConfig,
		Proxy:           http.ProxyFromEnvironment,
	})

	testdryrun, _ := cf.GlobalFlags["testhookdryrun"].(bool)
	if testdryrun {
		roundTripper = createTesthookLogger(true)
	} else {
		testhook, _ := cf.GlobalFlags["testhook"].(bool)
		if testhook {
			roundTripper = createTesthookLogger(false)
		}
	}

	clientConfig := &services.Config{
		Token:            getToken(),
		Host:             hostPort,
		OverrideHost:     overrideHostPort,
		Scheme:           scheme,
		Timeout:          getTimeoutForConfig(),
		ResponseHandlers: []services.ResponseHandler{&services.DefaultRetryResponseHandler{}},
		RoundTripper:     roundTripper,
		ClientVersion:    scloudVersion,
		TenantScoped:     tenantScoped,
		Region:           region,
	}

	result, err := sdk.NewClient(clientConfig)
	if err != nil {
		util.Fatal(err.Error())
	}
	return result
}

// Returns the api service client pointing to the New Client in the SDK.
func apiClient() *sdk.Client {
	// getTenantName() will prompt for tenant if none specified
	return apiClientWithTenant(getTenantName())
}

// Returns the api service client pointing to the New Client in the SDK using the specified tenant.
func apiClientWithTenant(tenant string) *sdk.Client {
	env := getEnvironment()

	result := newClient(&env.APIService)
	if result != nil {
		result.SetDefaultTenant(tenant)
	}

	return result
}

type testhookLogger struct {
	transport        http.RoundTripper
	cancelBeforeSend bool
}

func createTesthookLogger(cancelBeforeSend bool) http.RoundTripper {
	return &testhookLogger{transport: &http.Transport{}, cancelBeforeSend: cancelBeforeSend}
}

// RoundTrip implements http.RoundTripper when using testhook flag
func (out *testhookLogger) RoundTrip(request *http.Request) (*http.Response, error) {
	requestURLParts := strings.Split(strings.Trim(request.URL.String(), " "), "/")
	requestURL := strings.Join(requestURLParts[4:], "/")

	fmt.Printf("REQUEST URL:%v\n", requestURL)
	fmt.Printf("REQUEST BODY:%v\n", request.Body)

	//write to stream file for scloud test
	scloudTestCache := Abspath(".scloudTestOutput")
	os.Remove(scloudTestCache)

	f, err := os.Create(scloudTestCache)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("REQUEST URL:%v\n", requestURL))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	_, err = f.WriteString(fmt.Sprintf("REQUEST BODY:%v\n", request.Body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if out.cancelBeforeSend {
		fmt.Println(ScloudTesExecCancledError{}.Error())
		fmt.Print(ScloudTestDone)
		return nil, ScloudTesExecCancledError{}
	}

	response, err := out.transport.RoundTrip(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nRESPONSE:\n%v\n", response)
	_, err = f.WriteString(fmt.Sprintf("\nRESPONSE:\n%v\n", response))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Print(ScloudTestDone)
	return response, err
}

type ScloudTesExecCancledError struct {
}

func (ste ScloudTesExecCancledError) Error() string {
	return "Test mode enabled - request not sent"
}
