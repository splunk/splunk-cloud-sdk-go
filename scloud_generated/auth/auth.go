/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package auth

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/golang/glog"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cli/assets"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cli/fcache"
	"golang.org/x/crypto/ssh/terminal"
)

var options struct {
	env      string
	tenant   string
	username string
	password string
	noPrompt bool   // disable prompting for args
	authURL  string // scheme://host:port
	hostURL  string // scheme://host:port
	port     string
	insecure string // needs to be a string so we can test if the flag is set
	scheme   string
	certFile string
}

const SCloudHome = "SCLOUD_HOME"

var ctxCache *fcache.Cache
var settings *fcache.Cache

// Returns an absolute path. If the given path is not absolute it looks
// for the environment variable SCLOUD HOME and is joined with that.
// If the environment variable is not defined, the path is joined with
// the path to the home dir
func abspath(p string) string {
	if path.IsAbs(p) {
		return p
	}
	scloudHome, ok := os.LookupEnv(SCloudHome)
	var root string
	var err error
	if ok {
		root = scloudHome
	} else {
		root, err = homedir.Dir()
		if err != nil {
			fatal(err.Error())
		}
	}
	return path.Join(root, p)
}

// Returns the name of the selected environment.
func getEnvironmentName() string {
	if options.env != "" {
		return options.env
	}
	if envName, ok := settings.GetString("env"); ok {
		return envName
	}
	if options.noPrompt {
		fatal("no environment")
	}
	envName := "prod" // default
	options.env = envName
	return envName
}

func getEnvironment() *Environment {
	name := getEnvironmentName()
	env, err := GetEnvironment(name)
	if err != nil {
		fatal(err.Error())
	}
	return env
}

// Returns the selected username.
func getUsername() string {
	if options.username != "" {
		return options.username
	}
	if username, ok := settings.GetString("username"); ok {
		return username
	}
	if options.noPrompt {
		fatal("no username")
	}
	var username string
	fmt.Print("Username: ")
	if _, err := fmt.Scanln(&username); err != nil {
		fatal(err.Error())
	}
	options.username = username
	return username
}

func getpass() (string, error) {
	fmt.Print("Password: ")
	data, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Returns the selected password.
func getPassword() string {
	if options.password != "" {
		return options.password
	}
	if options.noPrompt {
		fatal("no password")
	}
	password, err := getpass()
	if err != nil {
		fatal(err.Error())
	}
	return password
}

// Returns the selected app profile.
func getProfile() (map[string]string, error) {
	name := getProfileName()
	profile, err := GetProfile(name)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// Returns the name of the selected app profile.
func getProfileName() string {
	return getEnvironment().Profile
}

// Returns the selected tenant name.
func getTenantName() string {
	if options.tenant != "" {
		return options.tenant
	}
	if tenant, ok := settings.GetString("tenant"); ok {
		return tenant
	}
	if options.noPrompt {
		fatal("no tenant")
	}
	var tenant string
	fmt.Print("Tenant: ")
	if _, err := fmt.Scanln(&tenant); err != nil {
		fatal(err.Error())
	}
	options.tenant = tenant
	return tenant
}

// Returns auth url from passed-in options or local settings.
// If auth_url is not specified, returns ""
func getAuthURL() string {
	return getOptionSettings(options.authURL, "auth-url")
}

// Returns host url from passed-in options or local settings.
// If host_url is not specified, returns ""
func getHostURL() string {
	return getOptionSettings(options.hostURL, "host-url")
}

// Returns scheme from passed-in options or local settings.
// If ca-cert is not specified, returns ""
func getCaCert() string {
	return getOptionSettings(options.certFile, "ca-cert")
}

// Check the flag options first, fall back on settings
func getOptionSettings(option string, setting string) string {
	if option != "" {
		return option
	}
	if setting, ok := settings.GetString(setting); ok {
		return setting
	}
	return ""
}

// Defaults to false, reads from settings first.
// Overridden by --insecure flag
func isInsecure() bool {
	insecure := false
	var err error
	// local settings cache default value
	if insecureStr, ok := settings.GetString("insecure"); ok {
		insecure, err = strconv.ParseBool(insecureStr)
		if err != nil {
			insecure = false
		}
	}
	// --insecure=true passed as global flag
	if options.insecure != "" {
		insecure, err = strconv.ParseBool(options.insecure)
		if err != nil {
			insecure = false
		}
	}
	if insecure {
		glog.Warningf("TLS certificate validation is disabled.")
	}
	return insecure
}

// Prints an error message to stderr.
func eprint(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "error: %s\n", msg)
}

func etoofew() {
	fatal("too few arguments")
}

// Prints an error message and exits.
func fatal(msg string, args ...interface{}) {
	eprint(msg, args...)
	os.Exit(1)
}

func parseArgs() []string {
	flag.StringVar(&options.env, "env", "", "environment name")
	flag.StringVar(&options.username, "u", "", "user name")
	flag.StringVar(&options.password, "p", "", "password")
	flag.StringVar(&options.tenant, "tenant", "", "tenant name")
	flag.StringVar(&options.authURL, "auth-url", "", "auth url")
	flag.StringVar(&options.hostURL, "host-url", "", "host url")
	flag.BoolVar(&options.noPrompt, "no-prompt", false, "disable prompting")
	flag.StringVar(&options.insecure, "insecure", "", "disable tls cert validation")
	flag.StringVar(&options.certFile, "ca-cert", "", "client certificate file")
	flag.Parse()
	return flag.Args()
}

// Verify that the given list is empty, or fatal.
func checkEmpty(items []string) {
	if len(items) > 0 {
		fatal("unexpected arguments: '%s'", strings.Join(items, ", "))
	}
}

// Ensure that the given app profile contains the required user credentials.
func ensureCredentials(profile map[string]string) {
	kind, ok := profile["kind"]
	if !ok {
		return
	}
	if kind == "client" {
		return // user creds not needed
	}
	if _, ok := profile["username"]; !ok {
		profile["username"] = getUsername()
	}
	if _, ok := profile["password"]; !ok {
		profile["password"] = getPassword()
	}
}

// Returns the cached authorization context associated with the given clientID.
func getCurrentContext(clientID string) *idp.Context {
	v := ctxCache.Get(clientID)
	m, ok := v.(*toml.Tree)
	if !ok {
		glog.Warningf("Deleting context cache")
		// bad cache entry
		ctxCache.Delete(clientID) //nolint:errcheck
		return nil
	}
	context := &idp.Context{}
	if err := FromToml(context, m); err != nil {
		eprint(err.Error())
		// bad cache entry
		ctxCache.Delete(clientID) //nolint:errcheck
		return nil
	}
	return context
}

// Returns an authorization "context", which consists of the OAuth token(s)
// and related metadata that correspond to a given app. If a valid cached
// context exists, return those, otherwise dispatch an authn flow that
// corresponds to the selected app profile.
func getContext() *idp.Context {
	profile, err := getProfile()
	if err != nil {
		fatal(err.Error())
		return nil
	}
	clientID, ok := profile["client_id"]
	if !ok {
		fatal("bad app profile: no client_id")
		return nil
	}
	context := getCurrentContext(clientID)
	if context != nil {
		// todo: re-authenticate if token has expired
		return context
	}
	ensureCredentials(profile)
	context, err = authenticate(profile)
	if err != nil {
		fatal(err.Error())
		return nil
	}
	ctxCache.Set(clientID, Map(context))
	return context
}

func getToken() string {
	return getContext().AccessToken
}

// Authenticate, using the selected app profile.
func Login(args []string) (*idp.Context, error) {
	loadConfigs()
	checkEmpty(args)
	name := getProfileName()
	profile, err := GetProfile(name)
	if err != nil {
		return nil, err
	}
	clientID := profile["client_id"]
	glog.CopyStandardLogTo("INFO")

	glog.Infof("Authenticate profile=%s clientID=%s", name, clientID)
	ensureCredentials(profile)
	context, err := authenticate(profile)
	if err != nil {
		return nil, err
	}
	ctxCache.Set(clientID, Map(context))
	return context, nil
}

// Load config and settings.
func loadConfigs() error {
	if err := loadConfig(); err != nil {
		return err
	}
	settings, _ = fcache.Load(abspath(".scloud"))
	ctxCache, _ = fcache.Load(abspath(".scloud_context"))
	return nil
}

// Load default config asset.
func loadConfig() error {
	file, err := assets.Open("default.yaml")
	if err != nil {
		return fmt.Errorf("err loading default.yaml: %s", err)
	}
	return Load(file)
}



type Service struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Scheme string `yaml:"scheme"`
}

type IdpService struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Scheme string `yaml:"scheme"`
	Server string `yaml:"server"`
}

type Environment struct {
	APIService Service    `yaml:"api-service"`
	AppService Service    `yaml:"app-service"`
	IdpService IdpService `yaml:"idp-service"`
	Profile    string     `yaml:"profile"`
}

type Cfg struct {
	Profiles     map[string]map[string]string `yaml:"profiles"`
	Environments map[string]*Environment      `yaml:"environments"`
}

var config Cfg

func Load(reader io.Reader) error {
	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&config); err != nil {
		return err
	}
	return nil
}

// Load the named config file.
func LoadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return Load(file)
}

func Environments() map[string]*Environment {
	return config.Environments
}

func GetEnvironment(name string) (*Environment, error) {
	env, ok := config.Environments[name]
	if !ok {
		return nil, fmt.Errorf("not found: '%s'", name)
	}
	return env, nil
}

// Returns the named application profile.
func GetProfile(name string) (map[string]string, error) {
	profile, ok := config.Profiles[name]
	if !ok {
		return nil, fmt.Errorf("not found: '%s'", name)
	}
	_, ok = profile["kind"] // ensure 'kind' exists
	if !ok {
		return nil, fmt.Errorf("missing kind")
	}
	return profile, nil
}

func Profiles() map[string]map[string]string {
	return config.Profiles
}
