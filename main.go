package main

import (
	"crypto/rand"
	"fmt"

	"github.com/hashicorp/vault/shamir"
)

// Example provided by chatGPT
func main() {
	// Generate a 256-bit secret key
	secretKey := generatePrivateKey()

	// Define the number of parts to split the key into and the minimum number of parts required to reconstruct the key
	totalParts, threshold := 50, 3

	// Split the key
	shares := splitMPC(secretKey, totalParts, threshold)

	// reconstructedKey, err := shamir.Combine(shares[:threshold])
	// if err != nil {
	// 	panic(err)
	// }

	// For demonstration purposes, reconstruct the key using 3 of the shares
	reconstructedKey, err := shamir.Combine([][]byte{shares[0], shares[3], shares[8]})
	if err != nil {
		panic(err)
	}

	// Print the reconstructed key
	fmt.Printf("Reconstructed Key: %x\n", reconstructedKey)
}

func generatePrivateKey() []byte {
	// Generate a 256-bit secret key
	secretKey := make([]byte, 32)
	if _, err := rand.Read(secretKey); err != nil {
		panic(err)
	}

	// secretKey, _ = hex.DecodeString("89011e19e534299e0f6b1d44492b8b3a0f39c81dc6feae8056435a99d473bd61")
	fmt.Printf("secret Key: %x\n", secretKey)

	return secretKey
}

func splitMPC(privateKey []byte, totalParts, threshold int) [][]byte {
	shares, err := shamir.Split(privateKey, totalParts, threshold)
	if err != nil {
		panic(err)
	}

	// Print the key shares
	for i, share := range shares {
		fmt.Printf("Share %d: %x\n", i+1, share)
	}

	return shares
}
