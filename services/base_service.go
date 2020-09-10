package services

import (
	"net/http"
	"net/url"
)

// BaseService provides the interface between client and services
type BaseService struct {
	Client IClient
}

//Client implemenents these methods to become a type of IClient used by the BaseService
type IClient interface {
	BuildURLFromPathParams(url.Values, string, string, interface{}) (url.URL, error)
	Get(requestParams RequestParams) (*http.Response, error)
	Post(requestParams RequestParams) (*http.Response, error)
	Patch(requestParams RequestParams) (*http.Response, error)
	Put(requestParams RequestParams) (*http.Response, error)
	Delete(requestParams RequestParams) (*http.Response, error)
}
