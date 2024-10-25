package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	_ "crypto/tls/fipsonly"
)

func main() {
	fmt.Println("Go Standard Library: crypto")
	fmt.Println()

	iterations := 10

	max := 100
	fmt.Printf("Generating %d random numbers between 0 and %d...\n", iterations, max)
	for i := 0; i < iterations; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
		if err != nil {
			handleError("error generating random number", err)
		}
		number := nBig.Int64()
		fmt.Printf("  %d\n", number)
	}
	fmt.Println()

	fmt.Printf("Generating %d prime numbers...\n", iterations)
	for i := 0; i < iterations; i++ {
		prime, err := rand.Prime(rand.Reader, 128)
		if err != nil {
			handleError("error generating prime number", err)
		}
		fmt.Printf("  %d\n", prime)
	}
	fmt.Println()

	// A MAC (Message Authentication Code) is a cryptographic hash function
	// that uses a secret key to generate a hash of a message.
	// The hash is used to verify that the message has not been tampered with.
	// The most common MAC algorithm is HMAC (Hash-based Message Authentication Code).
	// HMACs are used with hash functions like MD5 or SHA-256 to generate a MAC value.
	// The result of an HMAC function is called "hmac" or "tag"
	fmt.Println("Generating HMAC tags of a message and secret key using different hash algorithms...")
	message := "The quick brown fox jumps over the lazy dog"
	data := []byte(message)
	secret := "secret"
	key := []byte(secret)
	// Can use MD5 or SHA-256 (or other hash functions) to generate the hash
	md5Hash := hmac.New(md5.New, key)
	_, err := md5Hash.Write(data)
	if err != nil {
		handleError("error creating mac using MD5", err)
	}
	md5Tag := md5Hash.Sum(nil)
	sha256Hash := hmac.New(sha256.New, key)
	_, err = sha256Hash.Write(data)
	if err != nil {
		handleError("error creating mac using SHA256", err)
	}
	sha256Tag := sha256Hash.Sum(nil)
	// The tags are not human-readable, so convert them to a hex strings which will be:
	// - 32 characters long to represent the 16 bytes of the MD5 hash
	// - 64 characters long to represent the 32 bytes of the SHA-256 hash
	md5TagHex := hex.EncodeToString(md5Tag)
	fmt.Printf("Hex value of HMAC tag using MD5: %s\n", md5TagHex)
	sha256TagHex := hex.EncodeToString(sha256Tag)
	fmt.Printf("Hex value of HMAC tag using SHA256: %s\n", sha256TagHex)
	fmt.Println()
}

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
