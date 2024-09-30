package oegbay

type PageSchema struct {
	Signiture   string       `yaml:"signiture"`
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Theme       string       `yaml:"theme"`
	Content     string       `yaml:"content"`
	CreatedAt   string       `yaml:"created_at"`
	UpdatedAt   string       `yaml:"updated_at"`
	Pages       []PageSchema `yaml:"pages"`
}

type Schema struct {
	Version     string       `yaml:"version"`
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	CreatedAt   string       `yaml:"created_at"`
	UpdatedAt   string       `yaml:"updated_at"`
	Pages       []PageSchema `yaml:"pages"`
}

func ExtractBookFromSchema(schema *Schema) (*Book, error) {
	return &Book{
		Name:        schema.Name,
		Description: schema.Description,
		CreatedAt:   schema.CreatedAt,
		UpdatedAt:   schema.UpdatedAt,
	}, nil
}

func listPagesFromPageSchemas(parentSigniture string, pageSchemas []PageSchema, pages []Page) []Page {
	for _, page := range pageSchemas {
		pages = append(pages, Page{
			Signiture:       page.Signiture,
			Name:            page.Name,
			Description:     page.Description,
			Theme:           page.Theme,
			Content:         page.Content,
			ParentSigniture: parentSigniture,
			CreatedAt:       page.CreatedAt,
			UpdatedAt:       page.UpdatedAt,
		})

		pages = listPagesFromPageSchemas(page.Signiture, page.Pages, pages)
	}

	return pages
}

func ExtractPagesFromSchema(schema *Schema) ([]Page, error) {
	var pages []Page
	pages = listPagesFromPageSchemas("", schema.Pages, pages)
	return pages, nil
}

func ExtractPageFromSchema(schema *Schema, signiture string) (*Page, error) {
	// implement later

	return nil, nil
}
