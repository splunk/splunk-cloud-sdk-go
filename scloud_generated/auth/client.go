package auth

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const (
	defaultScheme = "https"
	defaultPort   = "443"
)

const (
	maxRetryCount = 6
)

type retryHandler struct{}

var sdkclient *sdk.Client

// Implements exponential backoff & retry on 429 and 504 response codes.
func (handler retryHandler) HandleResponse(client *services.BaseClient, request *services.Request, response *http.Response) (*http.Response, error) {
	var msg string

	switch response.StatusCode {
	case http.StatusTooManyRequests:
		msg = "Too many requests"
	case http.StatusGatewayTimeout:
		msg = "Gateway timeout"
	default:
		return response, nil
	}

	if request.NumAttempts > maxRetryCount {
		glog.Errorf("%s, giving up", msg)
		return response, nil
	}

	millis := ((1 << request.NumAttempts) * 500) + rand.Intn(250)
	backoff := time.Duration(millis) * time.Millisecond
	glog.Warningf("%s, backoff/retry (%v) ..", msg, backoff)
	time.Sleep(backoff)

	// reinitialize body, otherwise it will be empty (!?)
	body, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	request.Body = body
	return client.Do(request)
}

func GetClient() (*sdk.Client, error) {
	if sdkclient == nil {
		glog.CopyStandardLogTo("INFO")

		err := loadConfig()
		if err != nil {
			return nil, err
		}

		sdkclient := apiClient()

		//TODO: delete this example
		actions, err := sdkclient.ActionService.GetAction("crudemail_1568674368563")
		if err != nil {
			fmt.Println(err)
		}
		pprint(actions)

		return sdkclient, nil
	}

	return sdkclient, nil
}

// Returns a service client ( points to the new SDK Client) based on the given service config.
func newClient(svc *Service) *sdk.Client {
	var hostURL = getHostURL()

	serviceURL, err := url.Parse(hostURL)

	if err != nil {
		glog.Errorf("%s, is not a valid url", hostURL)
		return nil
	}

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

	clientConfig := &services.Config{
		Token:            getToken(),
		OverrideHost:     hostPort,
		Scheme:           scheme,
		Timeout:          10 * time.Second,
		ResponseHandlers: []services.ResponseHandler{&retryHandler{}},
		RoundTripper: util.NewCustomSdkTransport(&GlogWrapper{}, &http.Transport{
			TLSClientConfig: tlsConfig,
			Proxy:           http.ProxyFromEnvironment,
		})}

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

	svc := &env.APIService
	result := newClient(svc)
	result.SetDefaultTenant(tenant)

	return result
}

//TODO: delete this
func pprint(value interface{}) {
	if value == nil {
		return
	}
	switch vt := value.(type) {
	case string:
		fmt.Print(vt)
		if !strings.HasSuffix(vt, "\n") {
			fmt.Println()
		}
	default:
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "    ")
		err := encoder.Encode(value)
		if err != nil {
			fatal("json pprint error: %s", err.Error())
		}
	}
}
