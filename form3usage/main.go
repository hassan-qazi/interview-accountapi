package main

import (
	"example.com/form3"
	"fmt"
)

func main() {

	client := form3.NewClient(nil)

	/*

		gb := "GB"
		name := []string{"Hassan"}

		accountData := form3.AccountData{
			ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4d9",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Attributes: &form3.AccountAttributes{
				Country: &gb,
				Name:    name,
			},
		}

		account, err := client.Accounts.Create(&accountData)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(*account)
		}

	*/

	/*
		account, err := client.Accounts.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4d9")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(*account)
		}
	*/

	/*

		id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4b9"

		delerr := client.Accounts.Delete(id, "0")

		if delerr != nil {
			fmt.Println(delerr)
		} else {
			fmt.Println("Deleted %v", id)
		}

	*/

	accounts, err := client.Accounts.List()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(accounts)
	}

}
