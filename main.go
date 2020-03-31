package main

import (
	"github.com/gopherjs/gopherjs/js"
	"./hcl"
)

func main() {
	js.Module.Get("exports").Set("parse", hcl.Parse)
	js.Module.Get("exports").Set("stringify", hcl.Stringify)
}

