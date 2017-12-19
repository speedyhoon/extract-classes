# extract-classes
Extract CSS class names &amp; ids from a string

[![go report card](https://goreportcard.com/badge/github.com/speedyhoon/extract-classes)](https://goreportcard.com/report/github.com/speedyhoon/extract-classes) 

## Installation
```go get github.com/speedyhoon/extract-classes```

## Usage
```go
package main

import "github.com/speedyhoon/extractClasses"

func main(){
	css := "input:out-of-range #id-name#second.third::selection input:out-of-range::selection"
	println(extractClasses.Extract(css))
}
```
Returns: ```[]string{"#id-name", "#second", ".third"}```

## Licence
MIT License (MIT)

ported from [string-extract-class-names](https://github.com/codsen/string-extract-class-names)
