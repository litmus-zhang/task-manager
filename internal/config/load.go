package config

import (
	"encoding/json"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func NewConfig() (*Config, error) {

	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultPath := os.Getenv("VAULT_PATH")

	if vaultAddr == "" || vaultToken == "" || vaultPath == "" {
		log.Fatal("Vault address, token and path are required")
	}

	client, err := vault.NewClient(&vault.Config{
		Address: vaultAddr,
	})
	if err != nil {
		return nil, err
	}

	client.SetToken(vaultToken)

	secret, err := client.Logical().Read(vaultPath)
	if err != nil {
		return nil, err
	}

	if secret == nil || secret.Data == nil {
		log.Fatal("No data found at the specified Vault path")
	}

	// Extract the nested data
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatal("Invalid data format in Vault response")
	}

	var config Config
	err = mapToStruct(data, &config)
	if err != nil {
		return nil, err
	}

	log.Printf("Config: %+v", data)
	log.Printf("All Config: %+v", config)

	return &config, nil
}

func mapToStruct(data map[string]interface{}, result interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}
