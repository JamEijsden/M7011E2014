package main

import (
	"github.com/go-martini/martini"
)

/*
Kollade in martini lite. Värkar ganska trevligt.
för att få in martini i sitt golang så behöver
man hämta ner det.

Detta görs lättast med att skriva
go get github.com/go-martini/martini
från terminalen/ cmd
*/

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
