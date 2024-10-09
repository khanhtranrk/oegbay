package domain

import "time"

type Page struct {
	Signiture       string
	Name            string
	Description     string
	Theme           string
	Content         string
	ParentSigniture string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
