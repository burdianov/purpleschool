package main

import (
	"fmt"
	"password/account"
)

func main() {
	fmt.Println("__Password Manager__")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func findAccount() {}

func deleteAccount() {}

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

func createAccount() {
	login := promptData("Enter your login")
	password := promptData("Enter your password")
	url := promptData("Enter URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Incorrect URL or Login format")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")

	var res string
	fmt.Scanln(&res)

	return res
}
