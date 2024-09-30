package oegbay

import (
	"encoding/json"
)

type Load struct {
	EngineType string `json:"engine_type"`
}

type Book struct {
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type Page struct {
	Signiture       string
	Name            string
	Description     string
	Theme           string
	Content         string
	ParentSigniture string
	CreatedAt       string
	UpdatedAt       string
}

func UnmarshalLoad(load string) (*Load, error) {
	var ld Load

	err := json.Unmarshal([]byte(load), &ld)
	if err != nil {
		return nil, err
	}

	return &ld, nil
}
