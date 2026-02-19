package main

import (
	"github.com/flimzy/gopherlint/test/testdata_etc/unused_exported/lib"
)

func main() {
	lib.PublicFunc()
}
