package settle

import (
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/schema"
)

type Process struct {
}

func (p *Process) ReadSchema(load *Load) (*schema.BookSchema, error) {
	return nil, nil
}

func (p *Process) SaveSchema(load *Load, sch *schema.BookSchema) error {
	return nil
}

func (p *Process) ReadBook(load *Load) (*domain.Book, error) {
	return nil, nil
}

func (p *Process) CreateBook(load *Load, book *domain.Book) error {
	return nil
}

func (p *Process) UpdateBook(load *Load, book *domain.Book) error {
	return nil
}

func (p *Process) DeleteBook(load *Load) error {
	return nil
}

func (p *Process) ReadPageContent(load *Load, page *domain.Page) error {
	return nil
}

func (p *Process) CreatePage(load *Load, page *domain.Page) error {
	return nil
}

func (p *Process) UpdatePage(load *Load, page *domain.Page) error {
	return nil
}

func (p *Process) DeletePage(load *Load, page *domain.Page) error {
	return nil
}
