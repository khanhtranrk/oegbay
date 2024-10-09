package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPageSchema_Page(t *testing.T) {
	now := time.Now()
	ps := &PageSchema{
		Signiture:       "test-signiture",
		ParentSigniture: "parent-signiture",
		Name:            "Test Page",
		Description:     "This is a test page",
		Theme:           "default",
		Content:         "Test content",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	page := ps.Page()

	assert.Equal(t, ps.Signiture, page.Signiture)
	assert.Equal(t, ps.ParentSigniture, page.ParentSigniture)
	assert.Equal(t, ps.Name, page.Name)
	assert.Equal(t, ps.Description, page.Description)
	assert.Equal(t, ps.Theme, page.Theme)
	assert.Equal(t, ps.Content, page.Content)
	assert.Equal(t, ps.CreatedAt, page.CreatedAt)
	assert.Equal(t, ps.UpdatedAt, page.UpdatedAt)
}
