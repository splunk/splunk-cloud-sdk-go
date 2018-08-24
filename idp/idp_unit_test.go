// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package idp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const host = "https://myhost.net"

func TestNewDefaultClient(t *testing.T) {
	client := NewDefaultClient(host)
	assert.Equal(t, client.Host, host)
	assert.Equal(t, client.PathAuthn, defaultAuthnPath)
	assert.Equal(t, client.PathAuthorize, defaultAuthorizePath)
	assert.Equal(t, client.PathKeys, defaultKeysPath)
	assert.Equal(t, client.PathToken, defaultTokenPath)
}

func TestNewClient(t *testing.T) {
	client := NewClient(host, "custom/authn", "custom/authz", "custom/keys", "custom/token")
	assert.Equal(t, client.Host, host)
	assert.Equal(t, client.PathAuthn, "custom/authn")
	assert.Equal(t, client.PathAuthorize, "custom/authz")
	assert.Equal(t, client.PathKeys, "custom/keys")
	assert.Equal(t, client.PathToken, "custom/token")
}

func TestEncode(t *testing.T) {
	data := []byte(`{"event":{"time":123456789},"foo":"bar"}`)
	var obj interface{}
	err := json.Unmarshal(data, &obj)
	require.Nil(t, err)
	reader, err := encode(obj)
	require.NotNil(t, reader)
	require.Nil(t, err)

	var b bytes.Buffer
	n, err := reader.WriteTo(&b)
	assert.NotZero(t, n)
	assert.Nil(t, err)
	assert.Equal(t, string(data), b.String())
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
	require.Nil(t, err)
	assert.Equal(t, ctx.TokenType, "Bearer")
	assert.Equal(t, ctx.AccessToken, "my.access.token")
	assert.Equal(t, ctx.ExpiresIn, 3600)
	assert.Equal(t, ctx.Scope, "offline_access")
	assert.Equal(t, ctx.IDToken, "my.id.token")
	assert.Equal(t, ctx.RefreshToken, "my.refresh.token")
}
