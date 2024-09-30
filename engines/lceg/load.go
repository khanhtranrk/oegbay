package lceg

import (
	"encoding/json"
)

type Load struct {
	Path string `json:"path"`
}

func UnmarshalLoad(load string) (*Load, error) {
	var ld Load
	err := json.Unmarshal([]byte(load), &ld)
	if err != nil {
		return nil, err
	}

	return &ld, nil
}
