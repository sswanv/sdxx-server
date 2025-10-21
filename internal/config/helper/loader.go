package helper

import (
	"encoding/json"

	"github.com/dobyte/due/v2/config"
)

type JsonLoader func(string) ([]map[string]interface{}, error)

func DefaultJsonLoader(file string) ([]map[string]any, error) {
	bytes := config.Get(file).Bytes()
	var jsonData []map[string]any
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		return nil, err
	}
	return jsonData, nil
}
