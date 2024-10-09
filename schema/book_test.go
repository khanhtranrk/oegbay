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

func TestNewBookSchema(t *testing.T) {
	book := &domain.Book{
		Name:        "Test Book",
		Description: "Test Description",
	}

	bookSchema := schema.NewBookSchema(book)

	assert.Equal(t, setting.DefaultVersion, bookSchema.Version)
	assert.Equal(t, "Test Book", bookSchema.Name)
	assert.Equal(t, "Test Description", bookSchema.Description)
	assert.WithinDuration(t, time.Now(), bookSchema.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), bookSchema.UpdatedAt, time.Second)
}

func TestBookSchema_Book(t *testing.T) {
	bookSchema := &schema.BookSchema{
		Name:        "Test Book",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	book := bookSchema.Book()

	assert.Equal(t, "Test Book", book.Name)
	assert.Equal(t, "Test Description", book.Description)
	assert.Equal(t, bookSchema.CreatedAt, book.CreatedAt)
	assert.Equal(t, bookSchema.UpdatedAt, book.UpdatedAt)
}

func TestBookSchema_Update(t *testing.T) {
	bookSchema := &schema.BookSchema{
		Name:        "Old Name",
		Description: "Old Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	book := &domain.Book{
		Name:        "New Name",
		Description: "New Description",
	}

	bookSchema.Update(book)

	assert.Equal(t, "New Name", bookSchema.Name)
	assert.Equal(t, "New Description", bookSchema.Description)
	assert.WithinDuration(t, time.Now(), bookSchema.UpdatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), book.UpdatedAt, time.Second)
}

func TestBookSchema_ListPages(t *testing.T) {
	page := schema.PageSchema{
		Signiture:   uuid.New().String(),
		Name:        "Test Page",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	bookSchema := &schema.BookSchema{
		Pages: []schema.PageSchema{page},
	}

	pages := bookSchema.ListPages()

	assert.Len(t, pages, 1)
	assert.Equal(t, "Test Page", pages[0].Name)
	assert.Equal(t, "Test Description", pages[0].Description)
}

func TestBookSchema_GetPage(t *testing.T) {
	signature := uuid.New().String()
	page := schema.PageSchema{
		Signiture:   signature,
		Name:        "Test Page",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	bookSchema := &schema.BookSchema{
		Pages: []schema.PageSchema{page},
	}

	foundPage, err := bookSchema.GetPage(signature)

	assert.NoError(t, err)
	assert.Equal(t, "Test Page", foundPage.Name)
	assert.Equal(t, "Test Description", foundPage.Description)
}

func TestBookSchema_CreatePage(t *testing.T) {
	bookSchema := &schema.BookSchema{}

	page := &domain.Page{
		Name:        "New Page",
		Description: "New Description",
	}

	err := bookSchema.CreatePage(page)

	assert.NoError(t, err)
	assert.Len(t, bookSchema.Pages, 1)
	assert.Equal(t, "New Page", bookSchema.Pages[0].Name)
	assert.Equal(t, "New Description", bookSchema.Pages[0].Description)
	assert.NotEmpty(t, bookSchema.Pages[0].Signiture)
	assert.WithinDuration(t, time.Now(), bookSchema.Pages[0].CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), bookSchema.Pages[0].UpdatedAt, time.Second)
}

func TestBookSchema_UpdatePage(t *testing.T) {
	signature := uuid.New().String()
	page := schema.PageSchema{
		Signiture:   signature,
		Name:        "Old Page",
		Description: "Old Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	bookSchema := &schema.BookSchema{
		Pages: []schema.PageSchema{page},
	}

	updatedPage := &domain.Page{
		Signiture:   signature,
		Name:        "Updated Page",
		Description: "Updated Description",
	}

	err := bookSchema.UpdatePage(updatedPage)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Page", bookSchema.Pages[0].Name)
	assert.Equal(t, "Updated Description", bookSchema.Pages[0].Description)
	assert.WithinDuration(t, time.Now(), bookSchema.Pages[0].UpdatedAt, time.Second)
}

func TestBookSchema_DeletePage(t *testing.T) {
	signature := uuid.New().String()
	page := schema.PageSchema{
		Signiture:   signature,
		Name:        "Test Page",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	bookSchema := &schema.BookSchema{
		Pages: []schema.PageSchema{page},
	}

	err := bookSchema.DeletePage(signature)

	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), bookSchema.Pages[0].DeletedAt, time.Second)
}
