package settle

import (
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/schema"
)

type Settle struct {
	Process interface {
		ReadSchema(load *Load) (*schema.BookSchema, error)
		SaveSchema(load *Load, sch *schema.BookSchema) error

		CreateBook(load *Load, book *domain.Book) error
		DeleteBook(load *Load) error

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

func (s *Settle) Get(load string) (*domain.Book, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return nil, err
	}

	return sch.Book(), nil
}

func (s *Settle) Create(load string, book *domain.Book) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	if err := s.Process.CreateBook(ld, book); err != nil {
		return err
	}

	sch := schema.NewBookSchema(book)

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		s.Process.DeleteBook(ld)
		return err
	}

	return nil
}

func (s *Settle) Update(load string, book *domain.Book) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return err
	}

	sch.Update(book)

	if err := s.Process.SaveSchema(ld, sch); err != nil {
		return err
	}

	return nil
}

func (s *Settle) ListPages(load string) ([]domain.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := s.Process.ReadSchema(ld)
	if err != nil {
		return nil, err
	}

	pages := sch.ListPages()

	return pages, nil
}

func (s *Settle) GetPage(load string, signiture string) (*domain.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
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

func (s *Settle) CreatePage(load string, page *domain.Page) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
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

func (s *Settle) UpdatePage(load string, page *domain.Page) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
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

func (s *Settle) DeletePage(load string, signiture string) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
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
