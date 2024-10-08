package main

import (
	"fmt"

	"github.com/khanhtranrk/oegbay"
	"github.com/khanhtranrk/oegbay/engine/settle"
)

func main() {
	engines := map[string]oegbay.Engine{
		"LC": settle.New(),
	}
	engineBay := oegbay.New(engines)
	load := `{"engine_type":"LC","path":"./book_test"}`
	book, err := engineBay.ListPages(load)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(book)
}
