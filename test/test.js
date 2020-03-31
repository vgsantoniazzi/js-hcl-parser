var HCL = require("../dist/hcl.js")
var fs = require("fs")
var assert = require("assert");

function readFile(filename){
  return fs.readFileSync(`./test/schemas/${filename}`, 'utf8');
}

describe("HCL", function() {
  describe("#parse", function() {
    it("should match output JSON", function() {
      assert.equal(HCL.parse(readFile("input.hcl")), readFile("output.json"));
    });
  });

  describe("#stringify", function() {
    it("should match output HCL", function() {
      assert.equal(HCL.stringify(readFile("input.json")), readFile("output.hcl"));
    });
  });
});