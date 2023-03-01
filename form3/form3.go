package form3

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
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

	apiBaseUrl := os.Getenv("API_BASE_URL")

	if apiBaseUrl == "" {
		apiBaseUrl = defaultBaseURL
	}

	baseURL, _ := url.Parse(apiBaseUrl)

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
