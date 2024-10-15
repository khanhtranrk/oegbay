package oegbay_settle_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/engine/settle"
	"github.com/khanhtranrk/oegbay/setting"
	"github.com/stretchr/testify/assert"
)

const dir = "./document_test"

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
			"Path": dir,
		},
	)

	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// _load, _ := engineBay.MarshalLoad(load)

	// __load, _ := engineBay.UnmarshalLoad(_load)

	if err := engineBay.Create(load, &document); err != nil {
		assert.Fail(t, err.Error())
		return

	}

	// hard
	assert.DirExists(t, dir)
	assert.FileExists(t, filepath.Join(dir, setting.SchemaFile))

	// soft
	assert.Equal(t, "Test Document", document.Name)
	assert.Equal(t, "Test Description", document.Description)
	assert.NotNil(t, document.CreatedAt)
	assert.NotNil(t, document.UpdatedAt)

	cleanContext()
}
