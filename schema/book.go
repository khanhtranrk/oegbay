package schema

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/setting"
)

type BookSchema struct {
	Version     string       `yaml:"version"`
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	CreatedAt   string       `yaml:"created_at"`
	UpdatedAt   string       `yaml:"updated_at"`
	Pages       []PageSchema `yaml:"pages"`
}

func NewBookSchema(s *domain.Book) *BookSchema {
	now := time.Now()

	sch := &BookSchema{
		Version:     setting.DefaultVersion,
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   now.String(),
		UpdatedAt:   now.String(),
	}

	s.CreatedAt = now.String()
	s.UpdatedAt = now.String()

	return sch
}

func (s *BookSchema) Book() *domain.Book {
	return &domain.Book{
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func (s *BookSchema) Update(book *domain.Book) {
	now := time.Now()
	s.Name = book.Name
	s.Description = book.Description
	s.UpdatedAt = now.String()
	book.UpdatedAt = now.String()
}

func (s *BookSchema) ListPages() []domain.Page {
	var pages []domain.Page

	for _, pgs := range s.Pages {
		pages = append(pages, *pgs.Page())
	}

	return pages
}

func (s *BookSchema) GetPage(signiture string) (*domain.Page, error) {
	for _, pgs := range s.Pages {
		if pgs.Signiture == signiture {
			return pgs.Page(), nil
		}
	}

	return nil, fmt.Errorf("page with signature %s not found", signiture)
}

func (s *BookSchema) CreatePage(page *domain.Page) error {
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
		CreatedAt:       now.String(),
		UpdatedAt:       now.String(),
	})

	page.Signiture = signature
	page.CreatedAt = now.String()
	page.UpdatedAt = now.String()

	return nil
}

func (s *BookSchema) UpdatePage(page *oegbay.Page) error {
	now := time.Now()
	for index, pg := range s.Pages {
		if pg.Signiture == page.Signiture {
			s.Pages[index].Name = page.Name
			s.Pages[index].Description = page.Description
			s.Pages[index].Theme = page.Theme
			s.Pages[index].UpdatedAt = now.String()

			page.UpdatedAt = now.String()

			return nil
		}
	}

	return fmt.Errorf("page with signature %s not found", page.Signiture)
}

func (s *BookSchema) DeletePage(signiture string) error {
	for index, pg := range s.Pages {
		if pg.Signiture == signiture {
			s.Pages = append(s.Pages[:index], s.Pages[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("page with signature %s not found", signiture)
}
