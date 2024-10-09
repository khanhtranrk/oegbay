package oegbay_settle_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/engine/settle"
	"github.com/khanhtranrk/oegbay/setting"
	"github.com/stretchr/testify/assert"
)

const dir = "./book_test"

func cleanContext() error {
	if _, err := os.Stat(dir); err == nil {
		return os.RemoveAll(dir)
	} else if os.IsNotExist(err) {
		return nil
	}

	return nil
}

func TestCreate(t *testing.T) {
	cleanContext()

	engines := map[string]oegbay.Engine{
		"SETTLE": settle.New(),
	}

	engineBay := oegbay.New(engines)

	load := fmt.Sprintf(`{"engine_type":"SETTLE","path":"%s"}`, dir)

	book := domain.Book{
		Name:        "Test Book",
		Description: "Test Description",
	}

	if err := engineBay.Create(load, &book); err != nil {
		assert.Fail(t, err.Error())
		return

	}

	// hard
	assert.DirExists(t, dir)
	assert.FileExists(t, filepath.Join(dir, setting.SchemaFile))

	// soft
	assert.Equal(t, "Test Book", book.Name)
	assert.Equal(t, "Test Description", book.Description)
	assert.NotNil(t, book.CreatedAt)
	assert.NotNil(t, book.UpdatedAt)

	cleanContext()
}
