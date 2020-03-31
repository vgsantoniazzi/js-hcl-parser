# js-hcl-parser

## Overview

This package is used to generate Javascript parser for HCL (HashiCorp Configuration Language) from official repository https://github.com/hashicorp/hcl

## Installing

```sh
$ npm install js-hcl-parser
```

## Usage

See complete example under [examples](examples/)

```js
var HCL = require("js-hcl-parser")

const hclInput = `
scale {
  from = 72
  to = 24
}
`

const jsonInput = `
{
  "scale": {
    "from": 72,
    "to": 72
  }
}
`

console.log(HCL.parse(hclInput))

console.log(HCL.stringify(jsonInput))
```

### Building

```sh
$ go get -u github.com/hashicorp/hcl
$ go get -u github.com/gopherjs/gopherjs
$ gopherjs build . -o dist/hcl.js -m
```

### Testing

```sh
$ go test ./test
```

```js
$ npm test
```

## Contributing

I :heart: Open source!

[Follow github guides for forking a project](https://guides.github.com/activities/forking/)

[Follow github guides for contributing open source](https://guides.github.com/activities/contributing-to-open-source/#contributing)

[Squash pull request into a single commit](http://eli.thegreenplace.net/2014/02/19/squashing-github-pull-requests-into-a-single-commit/)

## License

js-hcl-parser is released under the [MIT license](http://opensource.org/licenses/MIT).