package main

import (
	"fmt"
	"password/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Password Manager__")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func findAccount(vault *account.Vault) {
	url := promptData("Enter URL")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Account not found")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Enter URL")
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("Account deleted")
	} else {
		color.Red("Account not found")
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Choose a variant:")
	fmt.Println("1. Create Account")
	fmt.Println("2. Find Account")
	fmt.Println("3. Delete Account")
	fmt.Println("4. Exit")

	fmt.Scan(&variant)
	return variant
}

func createAccount(vault *account.Vault) {
	login := promptData("Enter your login")
	password := promptData("Enter your password")
	url := promptData("Enter URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Incorrect URL or Login format")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")

	var res string
	fmt.Scanln(&res)

	return res
}
