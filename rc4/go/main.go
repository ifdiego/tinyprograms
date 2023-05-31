package main

import (
	"fmt"
	"os"
)

var N = 256 // 2^8

// key scheduling
func ksa(key string) []int {
	S := make([]int, N)
	for i := 0; i < N; i++ {
		S[i] = i
	}
	j := 0
	for i := 0; i < N; i++ {
		j = (j + S[i] + int(key[i%len(key)])) % N
		S[i], S[j] = S[j], S[i] // swap
	}
	return S
}

// pseudo random generation
func prga(S []int, n int) []int {
	j := 0
	keystream := []int{}
	for i := 0; i < n; i++ {
		temp := i
		i = (i + 1) % N
		j = (j + S[i]) % N
		S[i], S[j] = S[j], S[i]
		K := S[(S[i]+S[j])%N]
		keystream = append(keystream, K)
		i = temp
	}
	return keystream
}

func rc4(key string, plaintext string) []int {
	S := ksa(key)
	return prga(S, len(plaintext))
}

func main() {
	if len(os.Args) < 3 {
		panic(fmt.Sprintf("%s <key> <value>", os.Args[0]))
	}
	ciphertext := rc4(os.Args[1], os.Args[2])
	for index, char := range os.Args[2] {
		fmt.Printf("%02x", (int(char) ^ ciphertext[index]))
	}
	fmt.Println()
}
