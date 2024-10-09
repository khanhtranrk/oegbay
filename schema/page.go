package schema

import (
	"time"

	"github.com/khanhtranrk/oegbay/domain"
)

type PageSchema struct {
	Signiture       string    `yaml:"signiture"`
	ParentSigniture string    `yaml:"parent_signiture"`
	Name            string    `yaml:"name"`
	Description     string    `yaml:"description"`
	Theme           string    `yaml:"theme"`
	Content         string    `yaml:"content"`
	CreatedAt       time.Time `yaml:"created_at"`
	UpdatedAt       time.Time `yaml:"updated_at"`
	DeletedAt       time.Time `yaml:"deleted_at"`
}

func (ps *PageSchema) Page() *domain.Page {
	return &domain.Page{
		Signiture:       ps.Signiture,
		Name:            ps.Name,
		Description:     ps.Description,
		Theme:           ps.Theme,
		Content:         ps.Content,
		ParentSigniture: ps.ParentSigniture,
		CreatedAt:       ps.CreatedAt,
		UpdatedAt:       ps.UpdatedAt,
	}
}
