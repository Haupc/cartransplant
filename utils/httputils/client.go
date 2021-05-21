package httputils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpClient struct {
	params map[string]string
}

func NewHttpClient() *httpClient {
	return &httpClient{}
}

// AddParams ...
func (c *httpClient) AddParams(key, value string) {
	c.params[key] = value
}

// SetParams ...
func (c *httpClient) SetParams(params map[string]string) {
	c.params = params
}

// Get ...
func (c *httpClient) Get(url string) ([]byte, error) {
	for k, v := range c.params {
		url = fmt.Sprintf("%s&%s=%s", url, k, v)
	}
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(req.Body)
}
