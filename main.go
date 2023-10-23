package main

import (
	"fmt"
	"gocrypt/gocrypt"
)

func main() {
	plainText := "Hello World! My name is Den'ys"
	key := "123"

	cipherText := gocrypt.Encrypt_Pbox(plainText, key)
	falsh := "123"
	plain := gocrypt.Decrypt_Pbox(cipherText, falsh)
	fmt.Println("cipher -", cipherText)
	fmt.Println("plain -", plain)
}
