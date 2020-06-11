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
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/golang/glog"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth/fcache"
	cf "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/config"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v2"

	// Import needed to register files with fs
	_ "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth/statik"
)

const SCloudHome = "SCLOUD_HOME"
const DefaultEnv = "prod"

var ctxCache *fcache.Cache
var settings *fcache.Cache

// Returns an absolute path. If the given path is not absolute it looks
// for the environment variable SCLOUD HOME and is joined with that.
// If the environment variable is not defined, the path is joined with
// the path to the home dir
func Abspath(p string) string {
	if filepath.IsAbs(p) {
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
			util.Fatal(err.Error())
		}
	}
	return path.Join(root, p)
}

// Returns the name of the selected environment.
func GetEnvironmentName() string {
	err := loadConfigs()
	if err != nil {
		util.Fatal(err.Error())
	}

	usingEnv := DefaultEnv

	if envName, ok := settings.GetString("env"); ok && envName != "" {
		usingEnv = envName
	} else {
		util.Warning("No \"env\" is set in the config file, using default env instead")
	}

	util.Info("Using env - " + usingEnv)

	return usingEnv
}

func getEnvironment() *Environment {
	var name, env string
	env, _ = cf.GlobalFlags["env"].(string)

	if env != "" {
		name = env
	} else {
		name = GetEnvironmentName()
	}
	envName, err := GetEnvironment(name)
	if err != nil {
		util.Fatal(err.Error())
	}
	return envName
}

// Returns the selected username.
func getUsername() string {
	if username, ok := settings.GetString("username"); ok {
		return username
	}

	var username string
	fmt.Print("Username: ")
	if _, err := fmt.Scanln(&username); err != nil {
		util.Fatal(err.Error())
	}

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
func getPassword(cmd *cobra.Command) string {
	if cmd != nil {
		if pwd, err := cmd.Flags().GetString("pwd"); err == nil {
			if len(pwd) != 0 {
				return pwd
			}
		}
	}
	password, err := getpass()
	if err != nil {
		util.Fatal(err.Error())
	}

	return password
}

// Returns the selected app profile.
func getProfile() (map[string]string, error) {
	name := GetProfileName()
	profile, err := GetProfile(name)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// Returns the name of the selected app profile.
func GetProfileName() string {
	return getEnvironment().Profile
}

// Returns the selected tenant name.
func getTenantName() string {
	var tenant string

	tenant, _ = cf.GlobalFlags["tenant"].(string)
	if tenant != "" {
		return tenant
	}
	if tenant, ok := settings.GetString("tenant"); ok {
		return tenant
	}

	fmt.Print("Tenant: ")
	if _, err := fmt.Scanln(&tenant); err != nil {
		util.Fatal(err.Error())
	}

	return tenant
}

// Returns host url from passed-in options or local settings.
// If host_url is not specified, returns ""
func getHostURL() string {
	hostURL, _ := cf.GlobalFlags["host-url"].(string)
	if hostURL != "" {
		return hostURL
	}
	if setting, ok := settings.GetString("host-url"); ok {
		return setting
	}

	return ""
}

// Returns scheme from passed-in options or local settings.
// If ca-cert is not specified, returns ""
func getCaCert() string {
	cacert, _ := cf.GlobalFlags["ca-cert"].(string)
	if cacert != "" {
		return cacert
	}
	if setting, ok := settings.GetString("ca-cert"); ok {
		return setting
	}
	return ""
}

// Defaults to false, reads from settings first.
// Overridden by --insecure flag
func isInsecure() bool {
	insecure := false
	insecure, _ = cf.GlobalFlags["insecure"].(bool)
	if insecure != false {
		return insecure
	}
	// local settings cache default value
	if insecure, ok := settings.Get("insecure").(bool); ok {
		return insecure
	}

	return insecure
}

// Ensure that the given app profile contains the required user credentials.
func ensureCredentials(profile map[string]string, cmd *cobra.Command) {
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
		profile["password"] = getPassword(cmd)
	}
}

// Returns the cached authorization context associated with the given clientID.
func getCurrentContext(clientID string) *idp.Context {
	if ctxCache == nil {
		return nil
	}
	v := ctxCache.Get(clientID)
	m, ok := v.(*toml.Tree)
	if !ok {
		util.Warning("Deleting context cache")
		// bad cache entry
		ctxCache.Delete(clientID) //nolint:errcheck
		return nil
	}
	context := &idp.Context{}
	if err := FromToml(context, m); err != nil {
		util.Error(err.Error())
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
func getContext(cmd *cobra.Command) *idp.Context {
	profile, err := getProfile()
	if err != nil {
		util.Fatal(err.Error())
		return nil
	}
	clientID, ok := profile["client_id"]
	if !ok {
		util.Fatal("bad app profile: no client_id")
		return nil
	}
	context := getCurrentContext(clientID)
	if context != nil {
		// todo: re-authenticate if token has expired
		return context
	}

	kind, ok := profile["kind"]
	if !ok {
		util.Fatal("missing kind")
		return nil
	}

	authFlow, err := GetFlow(kind)
	if err != nil {
		util.Fatal(err.Error())
		return nil
	}

	context, err = authFlow(profile, cmd)
	if err != nil {
		util.Fatal(err.Error())
		return nil
	}

	ctxCache.Set(clientID, toMap(context))
	return context
}

func getToken() string {
	return getContext(nil).AccessToken
}

func LoginSetUp() error {
	return loadConfigs()
}

func GetEnvironmentProfile() (map[string]string, error) {
	name := GetProfileName()
	profile, err := GetProfile(name)
	if err != nil {
		return nil, err
	}
	return profile, err
}

// Authenticate, using the selected app profile.
func Login(cmd *cobra.Command, authFlow func(map[string]string, *cobra.Command) (*idp.Context, error)) (*idp.Context, error) {
	profile, err := GetEnvironmentProfile()
	if err != nil {
		return nil, err
	}

	clientID := profile["client_id"]
	glog.CopyStandardLogTo("INFO")

	context, err := authFlow(profile, cmd)
	if err != nil {
		return nil, err
	}
	ctxCache.Set(clientID, toMap(context))
	return context, nil
}

// Load config and settings.
func loadConfigs() error {
	if err := loadConfig(); err != nil {
		return err
	}

	var err error

	settings, err = fcache.Load(Abspath(viper.ConfigFileUsed()))

	if err != nil {
		return err
	}

	ctxCachePath := os.Getenv("SCLOUD_CACHE_PATH")

	if ctxCachePath == "" {
		ctxCachePath = ".scloud_context"
	}

	ctxCache, err = fcache.Load(Abspath(ctxCachePath))

	if err != nil {
		return err
	}

	return nil
}

// Load default config asset.
func loadConfig() error {
	file, err := open("default.yaml")
	if err != nil {
		return fmt.Errorf("err loading default.yaml: %s", err)
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return err
	}
	return nil
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

func GetEnvironment(name string) (*Environment, error) {
	env, ok := config.Environments[name]
	if !ok {
		return nil, fmt.Errorf("environment specified does not exist: '%s'", name)
	}
	return env, nil
}

// Returns the named application profile.
func GetProfile(name string) (map[string]string, error) {
	profile, ok := config.Profiles[name]
	if !ok {
		return nil, fmt.Errorf("auth.GetProfile key not found: '%s'", name)
	}
	_, ok = profile["kind"] // ensure 'kind' exists
	if !ok {
		return nil, fmt.Errorf("missing kind")
	}
	return profile, nil
}

// Returns the context from .scloud_context
func GetContext(cmd *cobra.Command) *idp.Context {
	return getContext(cmd)
}

// Set context in .scloud_context
func SetContext(cmd *cobra.Command, context map[string]interface{}) {
	profile, err := getProfile()
	if err != nil {
		util.Fatal(err.Error())
		return
	}

	clientID, ok := profile["client_id"]
	if !ok {
		util.Fatal("bad app profile: no client_id")
		return
	}
	ctxCache.Set(clientID, context)
}

// Open the named static file asset.
func open(fileName string) (io.Reader, error) {
	statikFs, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("assets.go: err calling fs.New() %v", err)
	}
	filePath := "/" + fileName
	httpFs := http.FileSystem(statikFs)
	b, err := fs.ReadFile(httpFs, filePath)
	if err != nil {
		return nil, fmt.Errorf("assets.go: err opening %s %v", filePath, err)
	}
	return bytes.NewReader(b), nil

}

func toMap(ctx *idp.Context) map[string]interface{} {
	result := map[string]interface{}{
		"token_type":   ctx.TokenType,
		"access_token": ctx.AccessToken,
		"expires_in":   ctx.ExpiresIn,
		"scope":        ctx.Scope}
	if ctx.IDToken != "" {
		result["id_token"] = ctx.IDToken
	}
	if ctx.RefreshToken != "" {
		result["refresh_token"] = ctx.RefreshToken
	}
	return result
}

// GlogWrapper is used to wrap glog.info() in a Print() function usable by splunk-cloud-sdk-go
type GlogWrapper struct {
}

func (gw *GlogWrapper) Print(v ...interface{}) {
	glog.Info(v...)
}
