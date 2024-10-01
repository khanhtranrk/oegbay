package lcop

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/engine/lceg"
	"github.com/khanhtranrk/oegbay/schema"
	"github.com/khanhtranrk/oegbay/setting"
	"gopkg.in/yaml.v3"
)

type LCOperation struct {
}

func GetSchema(load *lceg.Load) (*schema.BookSchema, error) {
	lp := filepath.Join(load.Path, setting.SchemaFile)
	data, err := os.ReadFile(lp)
	if err != nil {
		return nil, err
	}

	var sch schema.BookSchema
	if err := yaml.Unmarshal(data, &sch); err != nil {
		return nil, err
	}

	return &sch, nil
}

func SaveSchema(load *lceg.Load, sch *schema.BookSchema) error {
	data, err := yaml.Marshal(sch)
	if err != nil {
		return err
	}

	lp := filepath.Join(load.Path, setting.SchemaFile)
	return os.WriteFile(lp, data, 0755)
}

func Create(load *lceg.Load, book *oegbay.Book) error {
	info, err := os.Stat(load.Path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("path %s is a directory", load.Path)
	}

	return os.MkdirAll(load.Path, 0755)
}

func Delete(load *lceg.Load) error {
	return os.RemoveAll(load.Path)
}

func LoadPage(load *lceg.Load, page *oegbay.Page) error {
	lp := filepath.Join(load.Path, page.Signiture, setting.ContentFile)
	data, err := os.ReadFile(lp)
	if err != nil {
		return err
	}

	page.Content = string(data)
	return nil
}

func CreatePage(load *lceg.Load, page *oegbay.Page) error {
	fp := filepath.Join(load.Path, page.Signiture)
	if err := os.Mkdir(fp, 0755); err != nil {
		return err
	}

	lp := filepath.Join(load.Path, setting.ContentFile)
	if err := os.WriteFile(lp, []byte(page.Content), 0755); err != nil {
		os.RemoveAll(load.Path)
		return err
	}

	return nil
}

func UpdatePage(load *lceg.Load, page *oegbay.Page) error {
	lp := filepath.Join(load.Path, page.Signiture, setting.ContentFile)
	bp := filepath.Join(load.Path, page.Signiture, setting.ContentBackupFile)

	data, err := os.ReadFile(lp)
	if err != nil {
		return err
	}

	if err := os.WriteFile(bp, data, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(lp, []byte(page.Content), 0755); err != nil {
		return err
	}

	return nil
}

func DeletePage(load *lceg.Load, signiture string) error {
	lp := filepath.Join(load.Path, signiture)
	np := filepath.Join(load.Path, signiture+"_deleted")
	if err := os.Rename(lp, np); err != nil {
		return err
	}

	return nil
}
