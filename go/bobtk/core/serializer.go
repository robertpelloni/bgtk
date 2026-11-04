package core

import (
	"encoding/json"
	"fmt"
)

// Serializable represents an object capable of serialization (upp::Jsonize parity).
type Serializable interface {
	Serialize() (map[string]interface{}, error)
	Deserialize(data map[string]interface{}) error
}

// StateSerializer handles encoding and decoding UI state graphs and configuration trees.
type StateSerializer struct{}

func NewStateSerializer() *StateSerializer {
	return &StateSerializer{}
}

// ToJSON serializes any Go map or struct into a formatted JSON string.
func (s *StateSerializer) ToJSON(v interface{}) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to serialize state: %w", err)
	}
	return string(bytes), nil
}

// FromJSON deserializes a JSON string into target struct or map.
func (s *StateSerializer) FromJSON(jsonStr string, target interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), target); err != nil {
		return fmt.Errorf("failed to deserialize state: %w", err)
	}
	return nil
}
