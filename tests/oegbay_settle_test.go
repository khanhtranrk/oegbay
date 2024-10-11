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

	engines := map[string]oegbay.Engine{
		"SETTLE": settle.New(),
	}

	engineBay := oegbay.New(engines)

	document := oegbay.Document{
		Name:        "Test Document",
		Description: "Test Description",
	}

	load := &oegbay.Load{
		EngineType: "SETTLE",
		EngineLoad: &settle.Load{
			Path: dir,
		},
	}

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
