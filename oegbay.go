package oegbay

import "fmt"

type Engine interface {
	Get(load string) (*Book, error)
	Create(load string, book *Book) (*Book, error)
	Update(load string, book *Book) (*Book, error)
	Delete(load string) (*Book, error)

	ListPages(load string) ([]Page, error)
	GetPage(load string, signiture string) (*Page, error)
	CreatePage(load string, parentSigniture string, page *Page) (*Page, error)
	UpdatePage(load string, signiture string, page *Page) (*Page, error)
	DeletePage(load string, signiture string) (*Page, error)
}

type EngineBay struct {
	Engines map[string]Engine
}

func New(engines map[string]Engine) *EngineBay {
	return &EngineBay{
		Engines: engines,
	}
}

func (eb *EngineBay) Get(load string) (*Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].Get(load)
}

func (eb *EngineBay) Create(load string, book *Book) (*Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].Create(load, book)
}

func (eb *EngineBay) Update(load string, book *Book) (*Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].Update(load, book)
}

func (eb *EngineBay) Delete(load string) (*Book, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].Delete(load)
}

func (eb *EngineBay) ListPages(load string) ([]Page, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].ListPages(load)
}

func (eb *EngineBay) GetPage(load string, signiture string) (*Page, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].GetPage(load, signiture)
}

func (eb *EngineBay) CreatePage(load string, parentSigniture string, page *Page) (*Page, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].CreatePage(load, parentSigniture, page)
}

func (eb *EngineBay) UpdatePage(load string, parentSigniture string, page *Page) (*Page, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].UpdatePage(load, parentSigniture, page)
}

func (eb *EngineBay) DeletePage(load string, signiture string) (*Page, error) {
	ld, err := UnmarshalLoad(load)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal load: %v", err)
	}

	return eb.Engines[ld.EngineType].DeletePage(load, signiture)
}
