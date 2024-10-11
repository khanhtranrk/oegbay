package schema_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/schema"
	"github.com/khanhtranrk/oegbay/setting"
	"github.com/stretchr/testify/assert"
)

func TestNewDocumentSchema(t *testing.T) {
	document := &domain.Document{
		Name:        "Test Document",
		Description: "Test Description",
	}

	documentSchema := schema.NewDocumentSchema(document)

	assert.Equal(t, setting.DefaultVersion, documentSchema.Version)
	assert.Equal(t, "Test Document", documentSchema.Name)
	assert.Equal(t, "Test Description", documentSchema.Description)
	assert.WithinDuration(t, time.Now(), documentSchema.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), documentSchema.UpdatedAt, time.Second)
}

func TestDocumentSchema_Document(t *testing.T) {
	documentSchema := &schema.DocumentSchema{
		Name:        "Test Document",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	document := documentSchema.Document()

	assert.Equal(t, "Test Document", document.Name)
	assert.Equal(t, "Test Description", document.Description)
	assert.Equal(t, documentSchema.CreatedAt, document.CreatedAt)
	assert.Equal(t, documentSchema.UpdatedAt, document.UpdatedAt)
}

func TestDocumentSchema_Update(t *testing.T) {
	documentSchema := &schema.DocumentSchema{
		Name:        "Old Name",
		Description: "Old Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	document := &domain.Document{
		Name:        "New Name",
		Description: "New Description",
	}

	documentSchema.Update(document)

	assert.Equal(t, "New Name", documentSchema.Name)
	assert.Equal(t, "New Description", documentSchema.Description)
	assert.WithinDuration(t, time.Now(), documentSchema.UpdatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), document.UpdatedAt, time.Second)
}

func TestDocumentSchema_ListPages(t *testing.T) {
	page := schema.PageSchema{
		Signiture:   uuid.New().String(),
		Name:        "Test Page",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	documentSchema := &schema.DocumentSchema{
		Pages: []schema.PageSchema{page},
	}

	pages := documentSchema.ListPages()

	assert.Len(t, pages, 1)
	assert.Equal(t, "Test Page", pages[0].Name)
	assert.Equal(t, "Test Description", pages[0].Description)
}

func TestDocumentSchema_GetPage(t *testing.T) {
	signature := uuid.New().String()
	page := schema.PageSchema{
		Signiture:   signature,
		Name:        "Test Page",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	documentSchema := &schema.DocumentSchema{
		Pages: []schema.PageSchema{page},
	}

	foundPage, err := documentSchema.GetPage(signature)

	assert.NoError(t, err)
	assert.Equal(t, "Test Page", foundPage.Name)
	assert.Equal(t, "Test Description", foundPage.Description)
}

func TestDocumentSchema_CreatePage(t *testing.T) {
	documentSchema := &schema.DocumentSchema{}

	page := &domain.Page{
		Name:        "New Page",
		Description: "New Description",
	}

	err := documentSchema.CreatePage(page)

	assert.NoError(t, err)
	assert.Len(t, documentSchema.Pages, 1)
	assert.Equal(t, "New Page", documentSchema.Pages[0].Name)
	assert.Equal(t, "New Description", documentSchema.Pages[0].Description)
	assert.NotEmpty(t, documentSchema.Pages[0].Signiture)
	assert.WithinDuration(t, time.Now(), documentSchema.Pages[0].CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), documentSchema.Pages[0].UpdatedAt, time.Second)
}

func TestDocumentSchema_UpdatePage(t *testing.T) {
	signature := uuid.New().String()
	page := schema.PageSchema{
		Signiture:   signature,
		Name:        "Old Page",
		Description: "Old Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	documentSchema := &schema.DocumentSchema{
		Pages: []schema.PageSchema{page},
	}

	updatedPage := &domain.Page{
		Signiture:   signature,
		Name:        "Updated Page",
		Description: "Updated Description",
	}

	err := documentSchema.UpdatePage(updatedPage)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Page", documentSchema.Pages[0].Name)
	assert.Equal(t, "Updated Description", documentSchema.Pages[0].Description)
	assert.WithinDuration(t, time.Now(), documentSchema.Pages[0].UpdatedAt, time.Second)
}

func TestDocumentSchema_DeletePage(t *testing.T) {
	signature := uuid.New().String()
	page := schema.PageSchema{
		Signiture:   signature,
		Name:        "Test Page",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	documentSchema := &schema.DocumentSchema{
		Pages: []schema.PageSchema{page},
	}

	err := documentSchema.DeletePage(signature)

	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), documentSchema.Pages[0].DeletedAt, time.Second)
}
