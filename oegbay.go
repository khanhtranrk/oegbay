package oegbay

import (
	"github.com/khanhtranrk/oegbay/domain"
)

type Load = domain.Load
type Document = domain.Document
type Page = domain.Page

type Engine interface {
	Get(load interface{}) (*Document, error)
	Create(load interface{}, document *Document) error
	Update(load interface{}, document *Document) error

	ListPages(load interface{}) ([]domain.Page, error)
	GetPage(load interface{}, signiture string) (*Page, error)
	CreatePage(load interface{}, page *Page) error
	UpdatePage(load interface{}, page *Page) error
	DeletePage(load interface{}, signiture string) error
}

type EngineBay struct {
	Engines map[string]Engine
}

func New(engines map[string]Engine) *EngineBay {
	return &EngineBay{
		Engines: engines,
	}
}

func (eb *EngineBay) Get(load *Load) (*Document, error) {
	return eb.Engines[load.EngineType].Get(load.EngineLoad)
}

func (eb *EngineBay) Create(load *Load, document *Document) error {
	return eb.Engines[load.EngineType].Create(load.EngineLoad, document)
}

func (eb *EngineBay) Update(load *Load, document *Document) error {
	return eb.Engines[load.EngineType].Update(load.EngineLoad, document)
}

func (eb *EngineBay) ListPages(load *Load) ([]domain.Page, error) {
	return eb.Engines[load.EngineType].ListPages(load.EngineLoad)
}

func (eb *EngineBay) GetPage(load *Load, signiture string) (*Page, error) {
	return eb.Engines[load.EngineType].GetPage(load.EngineLoad, signiture)
}

func (eb *EngineBay) CreatePage(load *Load, parentSigniture string, page *Page) error {
	return eb.Engines[load.EngineType].CreatePage(load.EngineLoad, page)
}

func (eb *EngineBay) UpdatePage(load *Load, parentSigniture string, page *Page) error {
	return eb.Engines[load.EngineType].UpdatePage(load.EngineLoad, page)
}

func (eb *EngineBay) DeletePage(load *Load, signiture string) error {
	return eb.Engines[load.EngineType].DeletePage(load.EngineLoad, signiture)
}
