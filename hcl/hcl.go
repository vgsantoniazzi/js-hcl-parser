package hcl

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"

	hcl2 "github.com/hashicorp/hcl/v2"
	hcl2syntax "github.com/hashicorp/hcl/v2/hclsyntax"
)

func Parse(input string) string {
	var ast interface{}

	file, diags := hcl2syntax.ParseConfig([]byte(input), "file.hcl", hcl2.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		err := hcl.Unmarshal([]byte(input), &ast)
		if err != nil {
			return fmt.Sprintf("unable to parse JSON: %s", diags)
		}
	}
	ast, err := convertFile(file)
	if err != nil {
		err := hcl.Unmarshal([]byte(input), &ast)
		if err != nil {
			return fmt.Sprintf("unable to convert to JSON: %s", err)
		}
	}

	data, err := json.MarshalIndent(ast, "", "    ")
	if err != nil {
		return fmt.Sprintf("unable to print JSON: %s", err)
	}

	return string(data)
}

func Stringify(input string) string {
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
