package schema

import "github.com/khanhtranrk/oegbay"

type PageSchema struct {
	Signiture       string `yaml:"signiture"`
	ParentSigniture string `yaml:"parent_signiture"`
	Name            string `yaml:"name"`
	Description     string `yaml:"description"`
	Theme           string `yaml:"theme"`
	Content         string `yaml:"content"`
	CreatedAt       string `yaml:"created_at"`
	UpdatedAt       string `yaml:"updated_at"`
}

func (ps *PageSchema) Page() *oegbay.Page {
	return &oegbay.Page{
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
