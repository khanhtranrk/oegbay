package lceg

import (
	"os"
	"time"

	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/schema"
	"github.com/khanhtranrk/oegbay/setting"
)

type LCEngine struct {
	Operation interface {
		GetSchema(load *Load) (*schema.BookSchema, error)    // read schema file
		SaveSchema(load *Load, sch *schema.BookSchema) error // save schema file

		Create(load *Load, book *oegbay.Book) error // create book folder
		Delete(load *Load) error                    // delete book folder

		LoadPage(load *Load, page *oegbay.Page) error
		CreatePage(load *Load, page *oegbay.Page) error // create page folder and page files
		UpdatePage(load *Load, page *oegbay.Page) error // update page files
		DeletePage(load *Load, signiture string) error  // delete page folder
	}
}

func New() *LCEngine {
	return &LCEngine{}
}

func (lc *LCEngine) Get(load string) (*oegbay.Book, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return nil, err
	}

	return sch.Book(), nil
}

func (lc *LCEngine) Create(load string, book *oegbay.Book) (*oegbay.Book, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	if err := lc.Operation.Create(ld, book); err != nil {
		return nil, err
	}

	now := time.Now()
	sch := &schema.BookSchema{
		Version:     setting.DefaultVersion,
		Name:        book.Name,
		Description: book.Description,
		CreatedAt:   now.String(),
		UpdatedAt:   now.String(),
	}

	if err := lc.Operation.SaveSchema(ld, sch); err != nil {
		lc.Operation.Delete(ld)
		return nil, err
	}

	return book, nil
}

func (lc *LCEngine) Update(load string, book *oegbay.Book) (*oegbay.Book, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return nil, err
	}

	sch.Update(book)

	if err := lc.Operation.SaveSchema(ld, sch); err != nil {
		return nil, err
	}

	return book, nil
}

func (lc *LCEngine) Delete(load string) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	_, err = lc.Operation.GetSchema(ld)
	if err != nil {
		return err
	}

	if err := lc.Operation.Delete(ld); err != nil {
		return nil
	}

	os.RemoveAll(ld.Path)

	return nil
}

func (lc *LCEngine) ListPages(load string) ([]oegbay.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return nil, err
	}

	pages := sch.GetPages()

	return pages, nil
}

func (lc *LCEngine) GetPage(load string, signiture string) (*oegbay.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	sch, err := lc.Operation.GetSchema(ld)
	if err != nil {
		return nil, err
	}

	page, err := sch.GetPage(signiture)
	if err != nil {
		return nil, err
	}

	if err := lc.Operation.LoadPage(ld, page); err != nil {
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

func (lc *LCEngine) CreatePage(load string, page *oegbay.Page) (*oegbay.Page, error) {
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

func (lc *LCEngine) UpdatePage(load string, page *oegbay.Page) error {
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

func (lc *LCEngine) DeletePage(load string, signiture string) error {
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
