package main

import (
	"fmt"
	"goboxcrypt/goboxcrypt"
)

func main() {
	plainText := "Hello World and Distributed Lab team! My name is Den'ys. I wrote these encryption scripts. Look at README to understand how! :D"
	key := "3c45886ac87483a37dbc34ce7bf1fc99917a5ea8c9ba0ebab8a2d42a84d18c98"

	// encryption based on algo with Sbox swaps
	cipher_Sbox := goboxcrypt.Encrypt_Sbox(plainText, key)
	backtoPlane := goboxcrypt.Decrypt_Sbox(cipher_Sbox, key)
	fmt.Println("cipher -> ", cipher_Sbox)
	fmt.Println("plain text -> ", backtoPlane)
	/*
		cipher ->  b64f6e6e65c0e3658d6e68c0370668c08acad3b78dca83aab74f68c03d3783c0b74f374542c0ff21c00637454fc0cad3c08a4f069f21d36bc081c0fc8d65b74fc0b7804fd34fc04f061b8d2132b7ca6506c0d31b8dca32b7d36bc03d6565c4c037b7c0e2e4f78affe4c0b765c0aa06684f8dd3b7370668c08065fc42c07c8a
		plain text ->  Hello World and Distributed Lab team! My name is Den'ys. I wrote these encryption scripts. Look at README to understand how! :D
	*/

	// encryption based on algo with Pbox swaps, more interesting version
	cipher_Pbox := goboxcrypt.Encrypt_Pbox(plainText, key)
	backtoPlane = goboxcrypt.Decrypt_Pbox(cipher_Pbox, key)
	fmt.Println("cipher -> ", cipher_Pbox)
	fmt.Println("plain text -> ", backtoPlane)
	/*
		cipher ->  db5c5a5a504ff450775a5e4f5d525e4fde59757e7759577c7e5c5e4fda5d574f7e5c5d584d4fd8794f525d585c4f59754fde5c52447975424fd94f7477507e5c4f7e5b5c755c4f5c525577797f7e5950524f755577597f7e75424fda5050514f5d7e4ff7dcddded8dc4f7e504f7c525e5c77757e5d525e4f5b50744d4f63de
		plain text ->  Hello World and Distributed Lab team! My name is Den'ys. I wrote these encryption scripts. Look at README to understand how! :D
	*/
}
