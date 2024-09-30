package lceg

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/khanhtranrk/oegbay"
	"gopkg.in/yaml.v3"
)

type LCEngine struct {
}

func New() *LCEngine {
	return &LCEngine{}
}

func (lc *LCEngine) getSchema(load *Load) (*oegbay.Schema, error) {
	filePath := filepath.Join(load.Path, oegbay.InfoFile)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s: %w", filePath, err)
	}

	var schema oegbay.Schema
	err = yaml.Unmarshal(data, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data: %w", err)
	}

	return &schema, nil
}

func (lc *LCEngine) Get(load string) (*oegbay.Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	schema, err := lc.getSchema(ld)
	if err != nil {
		return nil, err
	}

	return oegbay.ExtractBookFromSchema(schema)
}

func (lc *LCEngine) Create(load string, book *oegbay.Book) (*oegbay.Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	if DirectoryExists(ld.Path) {
		return nil, fmt.Errorf("directory already exists at %s", ld.Path)
	}

	err = os.MkdirAll(ld.Path, 0755)
	if err != nil {
		return nil, err
	}

	schema := &oegbay.Schema{
		Version:     oegbay.DefaultVersion,
		Name:        book.Name,
		Description: book.Description,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	schemaData, err := yaml.Marshal(schema)
	if err != nil {
		os.RemoveAll(ld.Path)
		return nil, fmt.Errorf("failed to marshal YAML schema: %w", err)
	}

	filePath := filepath.Join(ld.Path, oegbay.InfoFile)
	infoFile, err := os.Create(filePath)
	if err != nil {
		os.RemoveAll(ld.Path)
		return nil, fmt.Errorf("failed to create file at %s: %w", filePath, err)
	}
	defer infoFile.Close()

	_, err = infoFile.Write(schemaData)
	if err != nil {
		os.RemoveAll(ld.Path)
		return nil, fmt.Errorf("failed to write to file at %s: %w", filePath, err)
	}

	return book, nil
}

func (lc *LCEngine) Update(load string, book *oegbay.Book) (*oegbay.Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	schema, err := lc.getSchema(ld)
	if err != nil {
		return nil, err
	}

	schema.Name = book.Name
	schema.Description = book.Description

	schemaData, err := yaml.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal YAML schema: %w", err)
	}

	filePath := filepath.Join(ld.Path, oegbay.InfoFile)
	err = os.WriteFile(filePath, schemaData, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to write to file at %s: %w", filePath, err)
	}

	return book, nil
}

func (lc *LCEngine) Delete(load string) (*oegbay.Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(ld.Path, oegbay.InfoFile)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s: %w", filePath, err)
	}

	var schema oegbay.Schema
	err = yaml.Unmarshal(data, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data: %w", err)
	}

	book, err := oegbay.ExtractBookFromSchema(&schema)
	if err != nil {
		return nil, fmt.Errorf("failed to extract book from schema: %w", err)
	}

	os.RemoveAll(ld.Path)

	return book, nil
}

func (lc *LCEngine) ListPages(load string) ([]oegbay.Page, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(ld.Path, oegbay.InfoFile)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s: %w", filePath, err)
	}

	var schema oegbay.Schema
	err = yaml.Unmarshal(data, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data: %w", err)
	}

	pages, err := oegbay.ExtractPagesFromSchema(&schema)
	if err != nil {
		return nil, fmt.Errorf("failed to extract page from schema: %w", err)
	}

	return pages, nil
}

func (lc *LCEngine) GetPage(load string, signiture string) (*oegbay.Page, error) {
	return nil, nil
}

func (lc *LCEngine) CreatePage(load string, parentSigniture string, page *oegbay.Page) (*oegbay.Page, error) {
	return nil, nil
}

func (lc *LCEngine) UpdatePage(load string, signiture string, page *oegbay.Page) (*oegbay.Page, error) {
	return nil, nil
}

func (lc *LCEngine) DeletePage(load string, signiture string) (*oegbay.Page, error) {
	return nil, nil
}
