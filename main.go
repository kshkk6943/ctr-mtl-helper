package main

import (
	"github.com/kshkk6943/ctr-mtl-helper/app"
)

func main() {
	csvTextReplace := app.NewCsvTextReplace()
	csvTextReplace.RunTextReplace("./")
}
