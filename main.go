package main

import (
	"fmt"
	"gocrypt/gocrypt"
)

func main() {
	plainText := "ia"

	cipherText := gocrypt.Encrypt_Pbox(plainText)
	plain := gocrypt.Decrypt_Pbox(cipherText)
	fmt.Println("cipher - ", cipherText)
	fmt.Println("plain - ", plain)
	//fmt.Println(gocrypt.HexToInt("96"))

	//fmt.Println(gocrypt.Permutate_P_box(97, "in"))
}
