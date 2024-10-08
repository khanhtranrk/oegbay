package oegbay

import (
	"encoding/json"

	"github.com/khanhtranrk/oegbay/domain"
)

type Engine interface {
	Get(load string) (*domain.Book, error)
	Create(load string, book *domain.Book) error
	Update(load string, book *domain.Book) error

	ListPages(load string) ([]domain.Page, error)
	GetPage(load string, signiture string) (*domain.Page, error)
	CreatePage(load string, page *domain.Page) error
	UpdatePage(load string, page *domain.Page) error
	DeletePage(load string, signiture string) error
}

type EngineBay struct {
	Engines map[string]Engine
}

func unmarshalLoad(load string) (*domain.Load, error) {
	var ld domain.Load
	if err := json.Unmarshal([]byte(load), &ld); err != nil {
		return nil, err
	}

	return &ld, nil
}

func New(engines map[string]Engine) *EngineBay {
	return &EngineBay{
		Engines: engines,
	}
}

func (eb *EngineBay) Get(load string) (*domain.Book, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	return eb.Engines[ld.EngineType].Get(load)
}

func (eb *EngineBay) Create(load string, book *domain.Book) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	return eb.Engines[ld.EngineType].Create(load, book)
}

func (eb *EngineBay) Update(load string, book *domain.Book) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	return eb.Engines[ld.EngineType].Update(load, book)
}

func (eb *EngineBay) ListPages(load string) ([]domain.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	return eb.Engines[ld.EngineType].ListPages(load)
}

func (eb *EngineBay) GetPage(load string, signiture string) (*domain.Page, error) {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return nil, err
	}

	return eb.Engines[ld.EngineType].GetPage(load, signiture)
}

func (eb *EngineBay) CreatePage(load string, parentSigniture string, page *domain.Page) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	return eb.Engines[ld.EngineType].CreatePage(load, page)
}

func (eb *EngineBay) UpdatePage(load string, parentSigniture string, page *domain.Page) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	return eb.Engines[ld.EngineType].UpdatePage(load, page)
}

func (eb *EngineBay) DeletePage(load string, signiture string) error {
	ld, err := unmarshalLoad(load)
	if err != nil {
		return err
	}

	return eb.Engines[ld.EngineType].DeletePage(load, signiture)
}
