package schema

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/setting"
)

type DocumentSchema struct {
	Version     string       `yaml:"version"`
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	CreatedAt   time.Time    `yaml:"created_at"`
	UpdatedAt   time.Time    `yaml:"updated_at"`
	Pages       []PageSchema `yaml:"pages"`
}

func NewDocumentSchema(s *domain.Document) *DocumentSchema {
	now := time.Now()

	sch := &DocumentSchema{
		Version:     setting.DefaultVersion,
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	s.CreatedAt = now
	s.UpdatedAt = now

	return sch
}

func (s *DocumentSchema) Document() *domain.Document {
	return &domain.Document{
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func (s *DocumentSchema) Update(document *domain.Document) {
	now := time.Now()
	s.Name = document.Name
	s.Description = document.Description
	s.UpdatedAt = now
	document.UpdatedAt = now
}

func (s *DocumentSchema) ListPages() []domain.Page {
	var pages []domain.Page

	for _, pgs := range s.Pages {
		pages = append(pages, *pgs.Page())
	}

	return pages
}

func (s *DocumentSchema) GetPage(signiture string) (*domain.Page, error) {
	for _, pgs := range s.Pages {
		if pgs.Signiture == signiture {
			return pgs.Page(), nil
		}
	}

	return nil, fmt.Errorf("page with signature %s not found", signiture)
}

func (s *DocumentSchema) CreatePage(page *domain.Page) error {
	now := time.Now()
	u := uuid.New().String()
	timestamp := now.Unix()
	signature := fmt.Sprintf("%s_%d", u, timestamp)

	s.Pages = append(s.Pages, PageSchema{
		Signiture:       signature,
		ParentSigniture: page.ParentSigniture,
		Name:            page.Name,
		Description:     page.Description,
		Theme:           page.Theme,
		Content:         page.Content,
		CreatedAt:       now,
		UpdatedAt:       now,
	})

	page.Signiture = signature
	page.CreatedAt = now
	page.UpdatedAt = now

	return nil
}

func (s *DocumentSchema) UpdatePage(page *domain.Page) error {
	now := time.Now()
	for index, pg := range s.Pages {
		if pg.Signiture == page.Signiture {
			s.Pages[index].Name = page.Name
			s.Pages[index].Description = page.Description
			s.Pages[index].Theme = page.Theme
			s.Pages[index].UpdatedAt = now

			page.UpdatedAt = now

			return nil
		}
	}

	return fmt.Errorf("page with signature %s not found", page.Signiture)
}

func (s *DocumentSchema) DeletePage(signiture string) error {
	for index, pg := range s.Pages {
		if pg.Signiture == signiture {
			s.Pages[index].DeletedAt = time.Now()
			return nil
		}
	}

	return fmt.Errorf("page with signature %s not found", signiture)
}
