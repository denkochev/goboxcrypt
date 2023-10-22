package main

import (
	"fmt"
	"gocrypt/gocrypt"
)

func main() {
	//gocrypt.Encrypt_S("hey")
	var str byte = "a"[0]
	fmt.Println("-----source-----")
	fmt.Printf("%b\n", str)
	fmt.Println("----------------")
	res := gocrypt.Permutate_P_box(str)
	fmt.Printf("%b\n", res)
	fmt.Println(res)
}
