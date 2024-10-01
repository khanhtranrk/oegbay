package schema

import (
	"testing"
	"time"

	"github.com/khanhtranrk/oegbay"
	"github.com/stretchr/testify/assert"
)

func TestBookSchema_Book(t *testing.T) {
	bookSchema := &BookSchema{
		Name:        "Test Book",
		Description: "Test Description",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	book := bookSchema.Book()

	assert.Equal(t, bookSchema.Name, book.Name)
	assert.Equal(t, bookSchema.Description, book.Description)
	assert.Equal(t, bookSchema.CreatedAt, book.CreatedAt)
	assert.Equal(t, bookSchema.UpdatedAt, book.UpdatedAt)
}

func TestBookSchema_Update(t *testing.T) {
	bookSchema := &BookSchema{
		Name:        "Test Book",
		Description: "Test Description",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	book := &oegbay.Book{
		Name:        "Updated Book",
		Description: "Updated Description",
	}

	bookSchema.Update(book)

	assert.Equal(t, book.Name, bookSchema.Name)
	assert.Equal(t, book.Description, bookSchema.Description)
	assert.NotEqual(t, book.CreatedAt, bookSchema.UpdatedAt)
	assert.Equal(t, book.UpdatedAt, bookSchema.UpdatedAt)
}

func TestBookSchema_GetPages(t *testing.T) {
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

	pages := bookSchema.GetPages()

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

	_, err = bookSchema.GetPage("nonexistent")
	assert.Error(t, err)
}

func TestBookSchema_CreatePage(t *testing.T) {
	bookSchema := &BookSchema{}
	page := &oegbay.Page{
		Name:        "New Page",
		Description: "New Description",
	}

	err := bookSchema.CreatePage(page)

	assert.NoError(t, err)
	assert.Len(t, bookSchema.Pages, 1)
	assert.Equal(t, "New Page", bookSchema.Pages[0].Name)
	assert.NotEmpty(t, bookSchema.Pages[0].Signiture)
	assert.NotEmpty(t, bookSchema.Pages[0].CreatedAt)
	assert.NotEmpty(t, bookSchema.Pages[0].UpdatedAt)
}

func TestBookSchema_UpdatePage(t *testing.T) {
	now := time.Now()
	bookSchema := &BookSchema{
		Pages: []PageSchema{
			{
				Signiture:   "page1",
				Name:        "Page 1",
				Description: "Description 1",
				CreatedAt:   now.String(),
				UpdatedAt:   now.String(),
			},
		},
	}

	page := &oegbay.Page{
		Signiture:   "page1",
		Name:        "Updated Page",
		Description: "Updated Description",
	}

	err := bookSchema.UpdatePage(page)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Page", bookSchema.Pages[0].Name)
	assert.Equal(t, "Updated Description", bookSchema.Pages[0].Description)
	assert.NotEqual(t, now.String(), bookSchema.Pages[0].UpdatedAt)
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
	assert.Len(t, bookSchema.Pages, 0)

	err = bookSchema.DeletePage("nonexistent")
	assert.Error(t, err)
}
