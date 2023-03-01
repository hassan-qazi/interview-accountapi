package form3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountsService service

type DataWrapper struct {
	Data *AccountData `json:"data,omitempty"`
}

type AccountListDataWrapper struct {
	Data *[]AccountData `json:"data,omitempty"`
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func (a AccountData) String() string {
	return fmt.Sprintf("Id: %v", a.ID)
}

type ApiError struct {
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode   int
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Response StatusCode: %v, ErrorMessage: %s",
		e.StatusCode, e.ErrorMessage)
}

func (s *AccountsService) List() (*[]AccountData, error) {

	url := "v1/organisation/accounts/"

	resp, err := s.Client.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = CheckAndBuildError(resp.StatusCode, &respBytes)

	if err != nil {
		return nil, err
	}

	respBody := &AccountListDataWrapper{}

	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (s *AccountsService) Create(accountData *AccountData) (*AccountData, error) {

	url := "v1/organisation/accounts"

	reqBody := &DataWrapper{
		Data: accountData,
	}

	resp, err := s.Client.Do("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = CheckAndBuildError(resp.StatusCode, &respBytes)

	if err != nil {
		return nil, err
	}

	respBody := &DataWrapper{}

	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (s *AccountsService) Fetch(id string) (*AccountData, error) {

	if id == "" {
		return nil, errors.New("id is empty")
	}

	url := fmt.Sprintf("v1/organisation/accounts/%v", id)

	resp, err := s.Client.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = CheckAndBuildError(resp.StatusCode, &respBytes)

	if err != nil {
		return nil, err
	}

	respBody := &DataWrapper{}

	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (s *AccountsService) Delete(id string, version *int64) error {

	if id == "" {
		return errors.New("id is empty")
	}

	url := fmt.Sprintf("v1/organisation/accounts/%v?version=%v", id, *version)

	resp, err := s.Client.Do("DELETE", url, nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = CheckAndBuildError(resp.StatusCode, &respBytes)

	if err != nil {
		return err
	}

	return nil
}

func CheckAndBuildError(statusCode int, respBytes *[]byte) error {

	if statusCode < 200 || statusCode > 299 {

		apiError := &ApiError{}

		apiError.StatusCode = statusCode
		apiError.ErrorMessage = http.StatusText(statusCode)

		if len(*respBytes) != 0 {

			err := json.Unmarshal(*respBytes, &apiError)
			if err != nil {
				return err
			}

			return apiError
		}

		return apiError
	}

	return nil
}
