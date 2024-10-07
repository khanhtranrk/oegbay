package locality

import (
	"encoding/json"
)

type Load struct {
	Path string `json:"path"`
}

func unmarshalLoad(load string) (*Load, error) {
	var ld Load
	if err := json.Unmarshal([]byte(load), &ld); err != nil {
		return nil, err
	}

	return &ld, nil
}
