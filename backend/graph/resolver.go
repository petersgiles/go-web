package graph

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct{}

// loadJSONData reads JSON data from the data directory
func loadJSONData(filename string, v interface{}) error {
	dataPath := filepath.Join("..", "data", filename)
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
