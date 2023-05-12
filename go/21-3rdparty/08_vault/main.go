package main

import (
	"context"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	fmt.Println("Vault")
	fmt.Println()

	config := vault.DefaultConfig()
	fmt.Printf("*** config=%v\n", config)

	config.Address = "http://127.0.0.1:8201" // Default port is 8200

	client, err := vault.NewClient(config)
	if err != nil {
		fmt.Printf("Error initializing client: %v", err)
	}

	// Authenticate
	// WARNING: This quickstart uses the root token for our Vault dev server.
	// Don't do this in production!
	//client.SetToken("dev-only-token")
	client.SetToken("hvs.5Pna7C23hcyTiLkfFMBw7b8l")

	secretData := map[string]interface{}{
		"password": "Hashi123",
	}

	ctx := context.Background()

	// Write a secret
	_, err = client.KVv2("secret").Put(ctx, "my-secret-password", secretData)
	if err != nil {
		fmt.Printf("Error writing secret: %v", err)
	}

	fmt.Println("Secret written successfully.")
}
