package integration

import (
	"example.com/form3"
	"github.com/google/uuid"
	"testing"
)

var client *form3.Client

func setup() {
	if client == nil {
		client = form3.NewClient(nil)
	}
}

func TestCreateAccount(t *testing.T) {

	setup()

	// create an account
	account, err := createAccount()
	if err != nil {
		t.Fatalf("createAccount returned error: %v", err)
	}

	t.Logf("Created account with ID: %v, successfully", account.ID)
}

func TestCreateAccountNilData(t *testing.T) {

	setup()

	_, err := client.Accounts.Create(nil)
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)
}

func TestCreateAccountMissingData(t *testing.T) {

	setup()

	gb := "GB"
	name := []string{"Hassan", "Qazi"}

	accountData := form3.AccountData{
		//ID:             (uuid.New()).String(),
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &form3.AccountAttributes{
			Country: &gb,
			Name:    name,
		},
	}

	_, err := client.Accounts.Create(&accountData)
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)
}

func TestCreateAccountDuplicateData(t *testing.T) {

	setup()

	gb := "GB"
	name := []string{"Hassan", "Qazi"}

	accountData := form3.AccountData{
		ID:             (uuid.New()).String(),
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &form3.AccountAttributes{
			Country: &gb,
			Name:    name,
		},
	}

	_, err := client.Accounts.Create(&accountData)
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	_, err = client.Accounts.Create(&accountData)
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)
}

func TestFetchAccount(t *testing.T) {

	setup()

	// create an account
	account, err := createAccount()
	if err != nil {
		t.Fatalf("createAccount returned error: %v", err)
	}

	t.Logf("Created account with ID: %v, successfully", account.ID)

	// fetch/read the account
	fetchedAccount, err := client.Accounts.Fetch(account.ID)
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}

	if fetchedAccount.ID != account.ID {
		t.Fatalf("Account ID should be %v, got %v", account.ID, fetchedAccount.ID)
	}

	t.Logf("Fetched account with ID: %v, successfully", fetchedAccount.ID)

}

func TestFetchAccountEmptyID(t *testing.T) {

	setup()

	// fetch/read the account
	_, err := client.Accounts.Fetch("")
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)

}

func TestFetchAccountUnknownID(t *testing.T) {

	setup()

	// create an account
	account, err := createAccount()
	if err != nil {
		t.Fatalf("createAccount returned error: %v", err)
	}

	t.Logf("Created account with ID: %v, successfully", account.ID)

	// fetch/read the account
	_, err = client.Accounts.Fetch((uuid.New()).String())
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)

}

func TestDeleteAccount(t *testing.T) {

	setup()

	// create an account
	account, err := createAccount()
	if err != nil {
		t.Fatalf("createAccount returned error: %v", err)
	}

	t.Logf("Created account with ID: %v, successfully", account.ID)

	// delete the account
	err = client.Accounts.Delete(account.ID, account.Version)
	if err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}

	// check if account exists
	fetchedAccount, err := client.Accounts.Fetch(account.ID)
	if fetchedAccount != nil && err == nil {
		t.Fatalf("Account with ID: %v not deleted", account.ID)

	}

	t.Logf("Deleted account with ID: %v, successfully", account.ID)

}

func TestDeleteAccountEmptyId(t *testing.T) {

	setup()

	version := int64(0)

	// delete the account
	err := client.Accounts.Delete("", &version)
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)

}

func TestDeleteAccountUnknownID(t *testing.T) {

	setup()

	version := int64(0)

	// delete the account
	err := client.Accounts.Delete((uuid.New()).String(), &version)
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)

}

func TestDeleteAccountUnknownVersion(t *testing.T) {

	setup()

	// create an account
	account, err := createAccount()
	if err != nil {
		t.Fatalf("createAccount returned error: %v", err)
	}

	t.Logf("Created account with ID: %v, successfully", account.ID)

	version := int64(1)

	// delete the account
	err = client.Accounts.Delete(account.ID, &version)
	if err == nil {
		t.Fatalf("Expecting err but got nil")
	}

	t.Logf("Error: %v", err)

}

func createAccount() (*form3.AccountData, error) {

	gb := "GB"
	name := []string{"Hassan", "Qazi"}

	accountData := form3.AccountData{
		ID:             (uuid.New()).String(),
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &form3.AccountAttributes{
			Country: &gb,
			Name:    name,
		},
	}

	account, err := client.Accounts.Create(&accountData)

	return account, err

}
