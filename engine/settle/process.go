package settle

import (
	"os"
	"path/filepath"

	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/schema"
	"github.com/khanhtranrk/oegbay/setting"
	"gopkg.in/yaml.v3"
)

type Process struct {
}

func (p *Process) ReadSchema(load *Load) (*schema.BookSchema, error) {
	data, err := os.ReadFile(filepath.Join(load.Path, setting.SchemaFile))
	if err != nil {
		return nil, err
	}

	var sch schema.BookSchema
	if err := yaml.Unmarshal(data, &sch); err != nil {
		return nil, err
	}

	return &sch, nil
}

func (p *Process) SaveSchema(load *Load, sch *schema.BookSchema) error {
	data, err := yaml.Marshal(sch)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(load.Path, setting.SchemaFile), data, 0755)
}

func (p *Process) CreateBook(load *Load, book *domain.Book) error {
	return os.MkdirAll(load.Path, 0755)
}

func (p *Process) DeleteBook(load *Load) error {
	return os.RemoveAll(load.Path)
}

func (p *Process) ReadPageContent(load *Load, page *domain.Page) error {
	data, err := os.ReadFile(filepath.Join(load.Path, page.Signiture, setting.ContentFile))
	if err != nil {
		return err
	}

	page.Content = string(data)

	return nil
}

func (p *Process) CreatePage(load *Load, page *domain.Page) error {
	err := os.MkdirAll(filepath.Join(load.Path, page.Signiture), 0755)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(load.Path, page.Signiture, setting.ContentFile), []byte(page.Content), 0755); err != nil {
		os.RemoveAll(filepath.Join(load.Path, page.Signiture))
		return err
	}

	return nil
}

func (p *Process) UpdatePage(load *Load, page *domain.Page) error {
	return os.WriteFile(filepath.Join(load.Path, page.Signiture, setting.ContentFile), []byte(page.Content), 0755)
}

func (p *Process) DeletePage(load *Load, page *domain.Page) error {
	return os.MkdirAll(filepath.Join(load.Path, page.Signiture), 0755)
}
