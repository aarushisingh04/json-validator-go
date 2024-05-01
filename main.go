package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func validateJSON(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	fmt.Println("JSON is valid")
	return nil
}

func main() {
	jsonFilePath := "example.json"
	err := validateJSON(jsonFilePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
