package oegbay_settle_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/engine/settle"
	"github.com/khanhtranrk/oegbay/setting"
	"github.com/stretchr/testify/assert"

	"github.com/otiai10/copy"
)

const documentTestPath = "./document_test"
const documentTestTemplatePath = "./resourses/settle/document_test"

func cleanContext() error {
	if _, err := os.Stat(documentTestPath); err == nil {
		return os.RemoveAll(documentTestPath)
	} else if os.IsNotExist(err) {
		return nil
	}

	return nil
}

func initContext() error {
	if err := copy.Copy(documentTestTemplatePath, documentTestPath); err != nil {
		return err
	}

	return nil
}

func TestCreate(t *testing.T) {
	cleanContext()

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	document := oegbay.Document{
		Name:        "Test Document",
		Description: "Test Description",
	}

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	if err := engineBay.Create(load, &document); err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// hard
	assert.DirExists(t, documentTestPath)
	assert.FileExists(t, filepath.Join(documentTestPath, setting.SchemaFile))

	// soft
	assert.Equal(t, "Test Document", document.Name)
	assert.Equal(t, "Test Description", document.Description)
	assert.NotNil(t, document.CreatedAt)
	assert.NotNil(t, document.UpdatedAt)

	cleanContext()
}

func TestGet(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	document, err := engineBay.Get(load)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "Document test", document.Name)
	assert.Equal(t, "Test document for testing the notebook schema.", document.Description)
	assert.NotNil(t, document.CreatedAt)
	assert.NotNil(t, document.UpdatedAt)

	cleanContext()
}

func TestUpdate(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	document := oegbay.Document{
		Name:        "Test Document",
		Description: "Test Description",
	}

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	if err := engineBay.Update(load, &document); err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// soft
	assert.Equal(t, "Test Document", document.Name)
	assert.Equal(t, "Test Description", document.Description)
	assert.NotNil(t, document.CreatedAt)
	assert.NotNil(t, document.UpdatedAt)

	cleanContext()
}

func TestListPage(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	pages, err := engineBay.ListPages(load)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Len(t, pages, 3)

	cleanContext()
}

func TestGetPage(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	page, err := engineBay.GetPage(load, "abaddf91-750f-40f3-b434-9bbe6081a6fb_1728967362")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "Test page", page.Name)
	assert.Equal(t, "This is the test page.", page.Description)
	assert.Equal(t, "#FF00FF", page.Theme)

	cleanContext()
}

func TestUpdatePage(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	page := oegbay.Page{
		Signiture:   "abaddf91-750f-40f3-b434-9bbe6081a6fb_1728967362",
		Name:        "Test Page",
		Description: "Test Description",
		Theme:       "#FF00FF",
	}

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	if err := engineBay.UpdatePage(load, &page); err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// soft
	assert.Equal(t, "Test Page", page.Name)
	assert.Equal(t, "Test Description", page.Description)
	assert.Equal(t, "#FF00FF", page.Theme)

	cleanContext()
}

func TestUpdatePageContent(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	if err := engineBay.UpdatePageContent(load, "abaddf91-750f-40f3-b434-9bbe6081a6fb_1728967362", []byte("This is the test page content.")); err != nil {
		assert.Fail(t, err.Error())
		return
	}

	cleanContext()
}

func TestDeletePage(t *testing.T) {
	if err := initContext(); err != nil {
		assert.Fail(t, err.Error())
	}

	engines := []oegbay.Engine{
		settle.New(),
	}

	engineBay := oegbay.New(engines)

	load, err := engineBay.NewLoadOfType(
		"settle",
		map[string]interface{}{
			"Path": documentTestPath,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	if err := engineBay.DeletePage(load, "abaddf91-750f-40f3-b434-9bbe6081a6fb_1728967362"); err != nil {
		assert.Fail(t, err.Error())
		return
	}

	cleanContext()
}
