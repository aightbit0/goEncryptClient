package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	gocrypt "github.com/aightbit0/goEncrypt"
)

func main() {

	for {
		var path string
		var password string
		var typeOfCrypt string
		var sure string

		fmt.Println("Welcome to goEncrypt")
		fmt.Print("Path to encrypt/decrypt Data -> ")

		scanner1 := bufio.NewScanner(os.Stdin)
		if scanner1.Scan() {
			path = scanner1.Text()
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("path not found")
		} else {
			fmt.Print("password -> ")
			scanner2 := bufio.NewScanner(os.Stdin)
			if scanner2.Scan() {
				password = scanner2.Text()
			}

			bytes := gocrypt.GenertateSecurePassword(password)
			key := hex.EncodeToString(bytes)
			//fmt.Printf("key to encrypt/decrypt : %s\n", key)
			fmt.Print("encrypt or decrypt -> ")
			scanner3 := bufio.NewScanner(os.Stdin)
			if scanner3.Scan() {
				typeOfCrypt = scanner3.Text()
			}

			allFiles := gocrypt.FillFiles(path, typeOfCrypt)

			for _, f := range allFiles {
				fmt.Println(f)
			}

			fmt.Print("are you sure to do this action y/n -> ")
			fmt.Scanln(&sure)

			if sure == "y" {
				if typeOfCrypt == "encrypt" {
					for _, v := range allFiles {
						gocrypt.Goencrypt(v, key)
					}
				} else if typeOfCrypt == "decrypt" {
					for _, v := range allFiles {
						gocrypt.Godecrypt(v, key)
					}
				} else {
					fmt.Println("Unknown command")
				}
			} else {
				fmt.Println("Cancelled")
			}
		}
	}
}
