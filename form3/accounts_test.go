package form3

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	gb     = "GB"
	name   = []string{"Hassan", "Qazi"}
	org_id = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
)

func TestAccountList(t *testing.T) {

	accountData := AccountData{
		ID:             (uuid.New()).String(),
		OrganisationID: org_id,
		Type:           "accounts",
		Attributes: &AccountAttributes{
			Country: &gb,
			Name:    name,
		},
	}

	accountList := []AccountData{accountData}

	respBody := &AccountListDataWrapper{
		Data: &accountList,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*respBody)

	}))
	defer svr.Close()

	c := NewClient(nil)

	baseURL, _ := url.Parse(svr.URL)

	c.BaseURL = baseURL

	if c.BaseURL.String() != svr.URL {
		t.Errorf("NewClient BaseURL is %v, want %v", c.BaseURL.String(), svr.URL)
	}

	list, err := c.Accounts.List()

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if accountData.ID != (*list)[0].ID {
		t.Errorf("Data Corrupted")
	}

}

func TestAccountCreate(t *testing.T) {

	accountData := AccountData{
		ID:             (uuid.New()).String(),
		OrganisationID: org_id,
		Type:           "accounts",
		Attributes: &AccountAttributes{
			Country: &gb,
			Name:    name,
		},
	}

	respBody := &DataWrapper{
		Data: &accountData,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*respBody)

	}))
	defer svr.Close()

	c := NewClient(nil)

	baseURL, _ := url.Parse(svr.URL)

	c.BaseURL = baseURL

	if c.BaseURL.String() != svr.URL {
		t.Errorf("NewClient BaseURL is %v, want %v", c.BaseURL.String(), svr.URL)
	}

	account, err := c.Accounts.Create(&accountData)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if accountData.ID != account.ID {
		t.Errorf("Data Corrupted")
	}

}

func TestAccountFetch(t *testing.T) {

	accountData := AccountData{
		ID:             (uuid.New()).String(),
		OrganisationID: org_id,
		Type:           "accounts",
		Attributes: &AccountAttributes{
			Country: &gb,
			Name:    name,
		},
	}

	respBody := &DataWrapper{
		Data: &accountData,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*respBody)

	}))
	defer svr.Close()

	c := NewClient(nil)

	baseURL, _ := url.Parse(svr.URL)

	c.BaseURL = baseURL

	if c.BaseURL.String() != svr.URL {
		t.Errorf("NewClient BaseURL is %v, want %v", c.BaseURL.String(), svr.URL)
	}

	account, err := c.Accounts.Fetch(accountData.ID)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if accountData.ID != account.ID {
		t.Errorf("Data Corrupted")
	}

}

func TestAccountDelete(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	}))
	defer svr.Close()

	c := NewClient(nil)

	baseURL, _ := url.Parse(svr.URL)

	c.BaseURL = baseURL

	if c.BaseURL.String() != svr.URL {
		t.Errorf("NewClient BaseURL is %v, want %v", c.BaseURL.String(), svr.URL)
	}

	id, version := (uuid.New()).String(), int64(0)

	err := c.Accounts.Delete(id, &version)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

}
