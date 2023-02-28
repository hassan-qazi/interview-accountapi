package form3

import (
	"bytes"
	"encoding/json"
	//"errors"
	//"fmt"
	"io"
	//"math/rand"
	"net/http"
	"net/url"
	//"strings"
	//"time"
)

const (
	defaultBaseURL   = "http://localhost:8080/"
	mediaType        = "application/json"
	defaultUserAgent = "go-form3"
)

type Client struct {
	HttpClient *http.Client
	BaseURL    *url.URL
	UserAgent  string
	Accounts   *AccountsService
	Common     service
}

type service struct {
	Client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{HttpClient: httpClient, BaseURL: baseURL, UserAgent: defaultUserAgent}

	c.Common.Client = c

	c.Accounts = (*AccountsService)(&c.Common)

	return c
}

func (c *Client) Do(method, urlStr string, body interface{}) (*http.Response, error) {

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", mediaType)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return c.HttpClient.Do(req)
}
