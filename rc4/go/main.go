package main

import (
	"fmt"
	"os"
)

func rc4(key string, plaintext string) string {
	return key + plaintext
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("errro")
	}
	ciphertext := rc4(os.Args[1], os.Args[2])
	fmt.Println("ciphertext=", ciphertext)
}
