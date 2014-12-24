package main

import (
	. "github.com/ttysteale/agenda/app"
	. "github.com/ttysteale/agenda/fs"
	"github.com/ttysteale/agenda/itemstore"
	"github.com/ttysteale/agenda/ui/widgets"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path := GetPath()
	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatal(err)
	}

	ui := widgets.Load()
	is := itemstore.NewItemStore(filepath.Join(path, "default.json"))
	app := NewApp(is, ui, 5)
	app.Load()
	ui.Run()
}
