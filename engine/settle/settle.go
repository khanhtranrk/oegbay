package settle

import (
	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/schema"
)

type Settle struct {
	Process interface {
		ReadSchema(load string) (*schema.BookSchema, error)
		SaveSchema(load string, sch *schema.BookSchema) error

		ReadBook(load string) (*domain.Book, error)
		CreateBook(load string, book *domain.Book) error
		DeleteBook(load string) error

		ReadPageContent(load string) ([]byte, error)
	}
}

// func New(operation operation.Operation) *Settle {
// 	return &Settle{
// 		Process: &Process{
// 			Operation: operation,
// 		},
// 	}
// }

func (s *Settle) Get(load string) (*domain.Book, error) {
	sch, err := s.Process.ReadSchema(load)
	if err != nil {
		return nil, err
	}

	return sch.Book(), nil
}

func (s *Settle) Create(load string, book *domain.Book) error {
	if err := s.Process.CreateBook(load, book); err != nil {
		return err
	}

	sch := schema.NewBookSchema(book)

	if err := s.Process.SaveSchema(load, sch); err != nil {
		s.Process.DeleteBook(load)
		return err
	}

	return nil
}

func (s *Settle) Update(load string, book *domain.Book) error {
	sch, err := s.Process.ReadSchema(load)
	if err != nil {
		return err
	}

	sch.Update(book)

	if err := s.Process.SaveSchema(load, sch); err != nil {
		return err
	}

	return nil
}

func (s *Settle) ListPages(load string) ([]domain.Page, error) {
	sch, err := s.Process.ReadSchema(load)
	if err != nil {
		return nil, err
	}

	pages := sch.ListPages()

	return pages, nil
}

func (s *Settle) GetPage(load string, signiture string) (*oegbay.Page, error) {
	sch, err := s.Process.ReadSchema(load)
	if err != nil {
		return nil, err
	}

	page, err := sch.GetPage(signiture)
	if err != nil {
		return nil, err
	}

	if err := s.Process.LoadPage(load, page); err != nil {
		return nil, err
	}

	// filePath := filepath.Join(ld.Path, signiture, setting.ContentFile)
	// content, err := os.ReadFile(filePath)
	// if err != nil {
	// 	return nil, err
	// }

	// page.Content = string(content)

	return page, nil
}

func (s *Settle) CreatePage(load string, page *oegbay.Page) (*oegbay.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return nil, err
	}

	if err := sch.CreatePage(page); err != nil {
		return nil, err
	}

	if err := lc.Operation.CreatePage(ld, page); err != nil {
		return nil, err
	}

	if err := lc.Operation.SaveSchema(ld, sch); err != nil {
		lc.Operation.DeletePage(ld, page.Signiture)
		return nil, err
	}

	return page, nil
}

func (s *Settle) UpdatePage(load string, page *oegbay.Page) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return err
	}

	if err := sch.UpdatePage(page); err != nil {
		return err
	}

	if err := lc.Operation.UpdatePage(ld, page); err != nil {
		return err
	}

	if err := lc.Operation.SaveSchema(ld, sch); err != nil {
		return err
	}

	// cntFilePath := filepath.Join(ld.Path, page.Signiture, setting.ContentFile)
	// bakFilePath := filepath.Join(ld.Path, page.Signiture, setting.ContentBackupFile)

	// content, err := os.ReadFile(cntFilePath)
	// if err != nil {
	// 	return err
	// }

	// err = os.WriteFile(bakFilePath, content, 0755)
	// if err != nil {
	// 	return err
	// }

	// err = os.WriteFile(cntFilePath, []byte(page.Content), 0755)
	// if err != nil {
	// 	return err
	// }

	// if err := lc.saveSchema(ld, sch); err != nil {
	// 	os.WriteFile(cntFilePath, content, 0755)
	// 	return err
	// }

	return nil
}

func (s *Settle) DeletePage(load string, signiture string) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return err
	}

	if err := sch.DeletePage(signiture); err != nil {
		return err
	}

	if err := lc.Operation.DeletePage(ld, signiture); err != nil {
		return err
	}

	if err := lc.Operation.SaveSchema(ld, sch); err != nil {
		return err
	}

	// pageFolderPath := filepath.Join(ld.Path, signiture)
	// pageDeletedFolderPath := filepath.Join(ld.Path, signiture+"_deleted")
	// err = os.Rename(pageFolderPath, pageDeletedFolderPath)
	// if err != nil {
	// 	return err
	// }

	// if err := lc.saveSchema(ld, sch); err != nil {
	// 	os.Rename(pageDeletedFolderPath, pageFolderPath)
	// 	return err
	// }

	return nil
}
