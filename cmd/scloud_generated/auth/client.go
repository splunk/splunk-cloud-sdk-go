package auth

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/golang/glog"
	cf "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/config"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const (
	defaultScheme = "https"
	defaultPort   = "443"
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
		err := loadConfig()
		if err != nil {
			return nil, err
		}

		sdkClient = apiClientWithTenant("system")

		return sdkClient, nil
	}

	return sdkClient, nil
}

// Returns a service client ( points to the new SDK Client) based on the given service config.
func newClient(svc *Service) *sdk.Client {
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

	host := serviceURL.Hostname()
	if host == "" {
		glog.Warningf("No host-url specified in config file, using default host")
		host = svc.Host
	}

	hostPort := host + ":" + port
	tlsConfig := &tls.Config{InsecureSkipVerify: isInsecure()}

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
			glog.Warningf("Failed to append %q to RootCAs: %v", caCert, err)
		}
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			glog.Warningf("No certs appended, using system certs only")
		}
		// set the RootCA
		tlsConfig.RootCAs = rootCAs
	}

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
		OverrideHost:     hostPort,
		Scheme:           scheme,
		Timeout:          10 * time.Second,
		ResponseHandlers: []services.ResponseHandler{&services.DefaultRetryResponseHandler{}},
		RoundTripper:     roundTripper,
	}

	result, err := sdk.NewClient(clientConfig)
	if err != nil {
		fatal(err.Error())
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

type testhooklogger struct {
	transport        http.RoundTripper
	cancelBeforeSend bool
}

func createTesthookLogger(cancelBeforeSend bool) http.RoundTripper {
	return &testhooklogger{transport: &http.Transport{}, cancelBeforeSend: cancelBeforeSend}
}

// RoundTrip implements http.RoundTripper when using testhook flag
func (out *testhooklogger) RoundTrip(request *http.Request) (*http.Response, error) {
	fmt.Printf("REQUEST URL:%v\n", request.URL)
	fmt.Printf("REQUEST BODY:%v\n", request.Body)

	if out.cancelBeforeSend {
		return nil, errors.New("For testrun, request was canceled")
	}

	response, err := out.transport.RoundTrip(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nRESPONSE:\n%v\n", response)

	return response, err
}
