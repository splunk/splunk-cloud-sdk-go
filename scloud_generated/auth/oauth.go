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
	"errors"
	"fmt"

	"github.com/pelletier/go-toml"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
)

// Returns the value corresponding to the given key, from the given map.
func gets(m map[string]string, k string) (string, error) {
	s, ok := m[k]
	if !ok {
		return "", fmt.Errorf("missing %s", k)
	}
	return s, nil
}

func getsd(m map[string]string, k, d string) (string, error) {
	s, ok := m[k]
	if !ok {
		return d, nil // default
	}
	return s, nil
}

func Map(ctx *idp.Context) map[string]interface{} {
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

func FromToml(ctx *idp.Context, t *toml.Tree) error {
	var ok bool
	var v interface{}
	if t.Has("token_type") {
		v = t.Get("token_type")
		if ctx.TokenType, ok = v.(string); !ok {
			return fmt.Errorf("bad value: token_type = %v", v)
		}
	}
	if t.Has("access_token") {
		v = t.Get("access_token")
		if ctx.AccessToken, ok = v.(string); !ok {
			return fmt.Errorf("bad value: access_token = %v", v)
		}
	}
	if t.Has("expires_in") {
		v = t.Get("expires_in")
		switch vt := v.(type) {
		case int:
			ctx.ExpiresIn = vt
		case float64:
			ctx.ExpiresIn = int(vt)
		case int64:
			ctx.ExpiresIn = int(vt)
		default:
			return fmt.Errorf("bad value: expires_in = %v (%T)", v, v)
		}
	}
	if t.Has("scope") {
		v = t.Get("scope")
		if ctx.Scope, ok = v.(string); !ok {
			return fmt.Errorf("bad value: scope = %v", v)
		}
	}
	if t.Has("id_token") {
		v = t.Get("id_token")
		if ctx.IDToken, ok = v.(string); !ok {
			return fmt.Errorf("bad value: id_token = %v", v)
		}
	}
	if t.Has("refresh_token") {
		v = t.Get("refresh_token")
		if ctx.RefreshToken, ok = v.(string); !ok {
			return fmt.Errorf("bad value: refresh_token = %v", v)
		}
	}
	return nil
}

func pkceFlow(profile map[string]string) (*idp.Context, error) {
	clientID, err := gets(profile, "client_id")
	if err != nil {
		return nil, err
	}
	redirectURI, err := gets(profile, "redirect_uri")
	if err != nil {
		return nil, err
	}
	scope, err := getsd(profile, "scope", "openid")
	if err != nil {
		return nil, err
	}
	username, err := gets(profile, "username")
	if err != nil {
		return nil, err
	}
	password, err := gets(profile, "password")
	if err != nil {
		return nil, err
	}
	idpHost, err := gets(profile, "idp_host")
	if err != nil {
		return nil, err
	}

	// Override idp_host from config file with -auth_url or auth_url in local settings
	authURL := getAuthURL()
	if authURL != "" {
		idpHost = authURL
	}

	tr := idp.NewPKCERetriever(clientID, redirectURI, idp.DefaultOIDCScopes, username, password, idpHost)

	// Allow on-prem to use insecure to bypass TLS Verification
	tr.Insecure = isInsecure()
	return tr.PKCEFlow(clientID, redirectURI, scope, username, password)
}

// Dispatch an authentication flow based on the given profile.
func authenticate(profile map[string]string) (*idp.Context, error) {
	kind, ok := profile["kind"]
	if !ok {
		return nil, errors.New("missing kind")
	}
	switch kind {
	case "pkce":
		return pkceFlow(profile)
	}
	return nil, fmt.Errorf("bad profile kind: '%s'", kind)
}
