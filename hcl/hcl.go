package hcl

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
)
func Parse(input string) (string) {
	var ast interface{}
	err := hcl.Unmarshal([]byte(input), &ast)
	if err != nil {
		return fmt.Sprintf("unable to parse JSON: %s", err)
	}

	data, err := json.MarshalIndent(ast, "", "    ")
	if err != nil {
		return fmt.Sprintf("unable to print JSON: %s", err)
	}

	return string(data)
}

func Stringify(input string) (string) {
	ast, err := hcl.Parse(input)
	if err != nil {
		return fmt.Sprintf("unable to parse HCL: %s", err)
	}

	var buf bytes.Buffer

	err = printer.Fprint(&buf, ast)
	if err != nil {
		return fmt.Sprintf("unable to print HCL: %s", err)
	}

	return buf.String()
}