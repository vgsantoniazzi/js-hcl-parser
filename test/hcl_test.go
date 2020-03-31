package hcl_test

import (
  "testing"
  "fmt"
  "io/ioutil"
  "../hcl"
)

func TestParse(t *testing.T){
  actual := hcl.Parse(file("input.hcl"))
  expected := file("output.json")
  if actual != expected {
    t.Errorf("\nexpected: %v\n Got: %v", expected, actual)
	}
}

func TestStringify(t *testing.T){
  actual := hcl.Stringify(file("input.json"))
  expected := file("output.hcl")
  if actual != expected {
    t.Errorf("\nexpected: %v\n Got: %v", expected, actual)
	}
}

func file(filename string) (string){
  b, err := ioutil.ReadFile(fmt.Sprintf("./schemas/%v", filename))
  if err != nil{
    return fmt.Sprintf("unable to print HCL: %s", err)
  }
  return string(b)
}
