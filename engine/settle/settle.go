package settle

import (
	"fmt"

	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/schema"
)

type Settle struct {
	Process interface {
		ReadSchema(load *Load) (*schema.DocumentSchema, error)
		SaveSchema(load *Load, sch *schema.DocumentSchema) error

		CreateDocument(load *Load, document *domain.Document) error
		DeleteDocument(load *Load) error

		ReadPageContent(load *Load, page *domain.Page) error
		CreatePage(load *Load, page *domain.Page) error
		UpdatePage(load *Load, page *domain.Page) error
		DeletePage(load *Load, page *domain.Page) error
	}
}

func New() *Settle {
	return &Settle{
		Process: &Process{},
	}
}

func (s *Settle) Get(load interface{}) (*domain.Document, error) {
	ld, ok := load.(*Load)
	if !ok {
		return nil, fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return nil, err
	}

	return sch.Document(), nil
}

func (s *Settle) Create(load interface{}, document *domain.Document) error {
	ld, ok := load.(*Load)
	if !ok {
		return fmt.Errorf("expected type *Load, got %T", load)
	}

	if err := s.Process.CreateDocument(ld, document); err != nil {
		return err
	}

	sch := schema.NewDocumentSchema(document)

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		s.Process.DeleteDocument(ld)
		return err
	}

	return nil
}

func (s *Settle) Update(load interface{}, document *domain.Document) error {
	ld, ok := load.(*Load)
	if !ok {
		return fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return err
	}

	sch.Update(document)

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		return err
	}

	return nil
}

func (s *Settle) ListPages(load interface{}) ([]domain.Page, error) {
	ld, ok := load.(*Load)
	if !ok {
		return nil, fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return nil, err
	}

	pages := sch.ListPages()

	return pages, nil
}

func (s *Settle) GetPage(load interface{}, signiture string) (*domain.Page, error) {
	ld, ok := load.(*Load)
	if !ok {
		return nil, fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return nil, err
	}

	page, err := sch.GetPage(signiture)
	if err != nil {
		return nil, err
	}

	if err := s.Process.ReadPageContent(ld, page); err != nil {
		return nil, err
	}

	return page, nil
}

func (s *Settle) CreatePage(load interface{}, page *domain.Page) error {
	ld, ok := load.(*Load)
	if !ok {
		return fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return err
	}

	if err := sch.CreatePage(page); err != nil {
		return err
	}

	if err := s.Process.CreatePage(ld, page); err != nil {
		return err
	}

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		s.Process.DeletePage(ld, page)
		return err
	}

	return nil
}

func (s *Settle) UpdatePage(load interface{}, page *domain.Page) error {
	ld, ok := load.(*Load)
	if !ok {
		return fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return err
	}

	if err := sch.UpdatePage(page); err != nil {
		return err
	}

	if err := s.Process.UpdatePage(ld, page); err != nil {
		return err
	}

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		return err
	}

	return nil
}

func (s *Settle) DeletePage(load interface{}, signiture string) error {
	ld, ok := load.(*Load)
	if !ok {
		return fmt.Errorf("expected type *Load, got %T", load)
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return err
	}

	if err := sch.DeletePage(signiture); err != nil {
		return err
	}

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		return err
	}

	return nil
}
