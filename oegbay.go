package oegbay

import (
	"encoding/json"
	"path"
	"reflect"
	"strings"

	"github.com/khanhtranrk/oegbay/domain"
)

type Load = domain.Load
type Document = domain.Document
type Page = domain.Page

type Engine interface {
	UnmarshalLoad(loadData interface{}) (interface{}, error)

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

func New(engines []Engine) *EngineBay {
	_engines := make(map[string]Engine)

	for _, engine := range engines {
		engineType := reflect.TypeOf(engine)
		if engineType.Kind() == reflect.Ptr {
			engineType = engineType.Elem()
		}
		_engines[strings.ToLower(engineType.Name())] = engine
	}

	return &EngineBay{
		Engines: _engines,
	}
}

func (eb *EngineBay) NewLoad(loadData interface{}) *Load {
	engineType := reflect.TypeOf(loadData)
	if engineType.Kind() == reflect.Ptr {
		engineType = engineType.Elem()
	}
	return &Load{
		EngineType: strings.ToLower(path.Base(engineType.PkgPath())),
		EngineLoad: loadData,
	}
}

func (eb *EngineBay) NewLoadOfType(engineType string, loadData interface{}) (*Load, error) {
	engineLoad, err := eb.Engines[engineType].UnmarshalLoad(loadData)
	if err != nil {
		return nil, err
	}

	return &Load{
		EngineType: engineType,
		EngineLoad: engineLoad,
	}, nil
}

func (eb *EngineBay) MarshalLoad(load *Load) ([]byte, error) {
	return json.Marshal(load)
}

func (eb *EngineBay) UnmarshalLoad(loadData []byte) (*Load, error) {
	var ld Load
	if err := json.Unmarshal(loadData, &ld); err != nil {
		return nil, nil
	}

	engineLoad, err := eb.Engines[ld.EngineType].UnmarshalLoad(ld.EngineLoad)
	if err != nil {
		return nil, err
	}

	ld.EngineLoad = engineLoad

	return &ld, nil
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
