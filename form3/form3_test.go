package form3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL is %v, want %v", c.BaseURL.String(), defaultBaseURL)
	}

	c2 := NewClient(nil)
	if c.HttpClient == c2.HttpClient {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}

func TestDo(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"data": []}`))
	}))
	defer svr.Close()

	c := NewClient(nil)

	baseURL, _ := url.Parse(svr.URL)

	c.BaseURL = baseURL

	if c.BaseURL.String() != svr.URL {
		t.Errorf("NewClient BaseURL is %v, want %v", c.BaseURL.String(), svr.URL)
	}

	resp, err := c.Do("GET", "v1/organisation/accounts/", nil)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	respBody := &AccountListDataWrapper{}

	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	dataLen := len(*(respBody.Data))

	if dataLen != 0 {
		t.Errorf("Expected empty list of AccountData got len: %v ", dataLen)
	}

}
