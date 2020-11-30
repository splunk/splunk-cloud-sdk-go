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

package idp

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	const providerHost = "https://myhost.net/"
	client := NewClient(providerHost, "custom/authn", "custom/authz", "custom/token", "custom/system/token", "custom/csrfToken", "custom/system/device", false)
	assert.Equal(t, client.ProviderHost, providerHost)
	assert.Equal(t, client.AuthnPath, "custom/authn")
	assert.Equal(t, client.AuthorizePath, "custom/authz")
	assert.Equal(t, client.TokenPath, "custom/token")
	assert.Equal(t, client.TenantTokenPath, "custom/system/token")
	assert.Equal(t, client.DevicePath, "custom/system/device")
	assert.Equal(t, client.CsrfTokenPath, "custom/csrfToken")
	clientEmptyParams := NewClient(providerHost, "", "", "", "", "", "", false)
	assert.Equal(t, clientEmptyParams.ProviderHost, providerHost)
	assert.Equal(t, clientEmptyParams.AuthnPath, defaultAuthnPath)
	assert.Equal(t, clientEmptyParams.AuthorizePath, defaultAuthorizePath)
	assert.Equal(t, clientEmptyParams.TokenPath, defaultTokenPath)
	assert.Equal(t, clientEmptyParams.TenantTokenPath, defaultTenantTokenPath)
	assert.Equal(t, clientEmptyParams.DevicePath, defaultDevicePath)
	assert.Equal(t, clientEmptyParams.CsrfTokenPath, defaultCsrfTokenPath)
}

// This is required to use bytes.Buffer in the http.Request.Body
type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func TestDecode(t *testing.T) {
	data := []byte(`{
		"token_type":"Bearer",
		"access_token":"my.access.token",
		"expires_in":3600,
		"scope":"offline_access",
		"id_token":"my.id.token",
		"refresh_token":"my.refresh.token"
	}`)
	resp := http.Response{
		StatusCode: 200,
		Body:       nopCloser{bytes.NewBuffer(data)},
	}
	ctx, err := decode(&resp)
	require.NotNil(t, ctx)
	require.NoError(t, err)
	assert.Equal(t, ctx.TokenType, "Bearer")
	assert.Equal(t, ctx.AccessToken, "my.access.token")
	assert.Equal(t, ctx.ExpiresIn, 3600)
	assert.Equal(t, ctx.Scope, "offline_access")
	assert.Equal(t, ctx.IDToken, "my.id.token")
	assert.Equal(t, ctx.RefreshToken, "my.refresh.token")
	assert.True(t, ctx.StartTime > 0)
}
