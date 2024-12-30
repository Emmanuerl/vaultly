package internal

import (
	"encoding/json"
)

// EncodeJSON takes in a value (preferably a struct with the right JSON tags)
// returns a JSON equivalent
func EncodeJSON(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// DecodeJSON converts a JSON encoded text and converts it based on the provided
// schem specified in v
func DecodeJSON(b []byte, v any) error {
	return json.Unmarshal(b, v)
}
