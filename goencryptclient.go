package main

import (
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

		fmt.Print("Path to encrypt/decrypt Data -> ")
		fmt.Scan(&path)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("path not found")
		} else {
			fmt.Print("password -> ")
			fmt.Scan(&password)
			bytes := gocrypt.GenertateSecurePassword(password)
			key := hex.EncodeToString(bytes)
			fmt.Printf("key to encrypt/decrypt : %s\n", key)
			fmt.Print("encrypt or decrypt -> ")
			fmt.Scan(&typeOfCrypt)
			allFiles := gocrypt.FillFiles(path, typeOfCrypt)

			for _, f := range allFiles {
				fmt.Println(f)
			}

			fmt.Print("are you sure to do this action y/n -> ")
			fmt.Scan(&sure)

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
