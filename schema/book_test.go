package schema

import (
	"testing"
	"time"

	"github.com/khanhtranrk/oegbay/domain"
	"github.com/khanhtranrk/oegbay/setting"
	"github.com/stretchr/testify/assert"
)

func TestNewBookSchema(t *testing.T) {
	book := &domain.Book{
		Name:        "Test Book",
		Description: "Test Description",
	}

	bookSchema := NewBookSchema(book)

	assert.Equal(t, setting.DefaultVersion, bookSchema.Version)
	assert.Equal(t, "Test Book", bookSchema.Name)
	assert.Equal(t, "Test Description", bookSchema.Description)
	assert.NotEmpty(t, bookSchema.CreatedAt)
	assert.NotEmpty(t, bookSchema.UpdatedAt)
}

func TestBookSchema_Book(t *testing.T) {
	bookSchema := &BookSchema{
		Name:        "Test Book",
		Description: "Test Description",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	book := bookSchema.Book()

	assert.Equal(t, "Test Book", book.Name)
	assert.Equal(t, "Test Description", book.Description)
	assert.Equal(t, bookSchema.CreatedAt, book.CreatedAt)
	assert.Equal(t, bookSchema.UpdatedAt, book.UpdatedAt)
}

func TestBookSchema_Update(t *testing.T) {
	bookSchema := &BookSchema{
		Name:        "Old Name",
		Description: "Old Description",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	book := &domain.Book{
		Name:        "New Name",
		Description: "New Description",
	}

	bookSchema.Update(book)

	assert.Equal(t, "New Name", bookSchema.Name)
	assert.Equal(t, "New Description", bookSchema.Description)
	assert.NotEqual(t, bookSchema.CreatedAt, bookSchema.UpdatedAt)
}

func TestBookSchema_ListPages(t *testing.T) {
	bookSchema := &BookSchema{
		Pages: []PageSchema{
			{
				Signiture:   "page1",
				Name:        "Page 1",
				Description: "Description 1",
			},
			{
				Signiture:   "page2",
				Name:        "Page 2",
				Description: "Description 2",
			},
		},
	}

	pages := bookSchema.ListPages()

	assert.Len(t, pages, 2)
	assert.Equal(t, "Page 1", pages[0].Name)
	assert.Equal(t, "Page 2", pages[1].Name)
}

func TestBookSchema_GetPage(t *testing.T) {
	bookSchema := &BookSchema{
		Pages: []PageSchema{
			{
				Signiture:   "page1",
				Name:        "Page 1",
				Description: "Description 1",
			},
		},
	}

	page, err := bookSchema.GetPage("page1")
	assert.NoError(t, err)
	assert.Equal(t, "Page 1", page.Name)

	_, err = bookSchema.GetPage("page2")
	assert.Error(t, err)
}

func TestBookSchema_CreatePage(t *testing.T) {
	bookSchema := &BookSchema{}

	page := &domain.Page{
		Name:        "New Page",
		Description: "New Description",
	}

	err := bookSchema.CreatePage(page)
	assert.NoError(t, err)
	assert.Len(t, bookSchema.Pages, 1)
	assert.Equal(t, "New Page", bookSchema.Pages[0].Name)
	assert.NotEmpty(t, bookSchema.Pages[0].Signiture)
}

func TestBookSchema_UpdatePage(t *testing.T) {
	bookSchema := &BookSchema{
		Pages: []PageSchema{
			{
				Signiture:   "page1",
				Name:        "Old Page",
				Description: "Old Description",
			},
		},
	}

	page := &domain.Page{
		Signiture:   "page1",
		Name:        "Updated Page",
		Description: "Updated Description",
	}

	err := bookSchema.UpdatePage(page)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Page", bookSchema.Pages[0].Name)
	assert.Equal(t, "Updated Description", bookSchema.Pages[0].Description)
}

func TestBookSchema_DeletePage(t *testing.T) {
	bookSchema := &BookSchema{
		Pages: []PageSchema{
			{
				Signiture:   "page1",
				Name:        "Page 1",
				Description: "Description 1",
			},
		},
	}

	err := bookSchema.DeletePage("page1")
	assert.NoError(t, err)
	assert.NotEmpty(t, bookSchema.Pages[0].DeletedAt)

	err = bookSchema.DeletePage("page2")
	assert.Error(t, err)
}
