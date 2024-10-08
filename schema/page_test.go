package schema

import (
	"testing"

	"github.com/khanhtranrk/oegbay/domain"
	"github.com/stretchr/testify/assert"
)

func TestPageSchema_Page(t *testing.T) {
	ps := &PageSchema{
		Signiture:       "test-signiture",
		ParentSigniture: "test-parent-signiture",
		Name:            "test-name",
		Description:     "test-description",
		Theme:           "test-theme",
		Content:         "test-content",
		CreatedAt:       "2023-01-01T00:00:00Z",
		UpdatedAt:       "2023-01-02T00:00:00Z",
		DeletedAt:       "2023-01-03T00:00:00Z",
	}

	page := ps.Page()

	expectedPage := &domain.Page{
		Signiture:       "test-signiture",
		ParentSigniture: "test-parent-signiture",
		Name:            "test-name",
		Description:     "test-description",
		Theme:           "test-theme",
		Content:         "test-content",
		CreatedAt:       "2023-01-01T00:00:00Z",
		UpdatedAt:       "2023-01-02T00:00:00Z",
	}

	assert.Equal(t, expectedPage, page)
}
