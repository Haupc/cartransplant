package httputils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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
func (c *httpClient) Get(_url string) ([]byte, error) {
	if !strings.Contains(_url, "?") && len(c.params) > 0 {
		_url = _url + "?"
	}
	for k, v := range c.params {
		_url = fmt.Sprintf("%s&%s=%s", _url, k, url.QueryEscape(v))
	}
	_url = strings.Replace(_url, "?&", "?", 1)
	log.Println(_url)
	req, err := http.Get(_url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(req.Body)
}
