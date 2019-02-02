# strmet

[![GoDoc](https://godoc.org/github.com/eskriett/strmet?status.svg)](https://godoc.org/github.com/eskriett/strmet)
[![Build Status](https://travis-ci.org/eskriett/strmet.svg?branch=master)](https://travis-ci.org/eskriett/strmet)

Fast and memory efficient string metric algorithms.

Available algorithms:
* [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)
* [Damerau–Levenshtein distance](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)

## Example

```go
package main

import (
    "fmt"
    "github.com/eskriett/strmet"
)

func main() {
    s1 := "baseball"
    s2 := "football"

    fmt.Printf("The Levenshtein distance between %s and %s is %d\n",
        s1, s2, strmet.Levenshtein(s1, s2, 10))
	// -> The Levenshtein distance between baseball and football is 4

    s1 = "town"
    s2 = "twon"
    fmt.Printf("The Damerau–Levenshtein distance between %s and %s is %d\n",
        s1, s2, strmet.DamerauLevenshtein(s1, s2, 10))
	// -> The Damerau–Levenshtein distance between town and twon is 1
}
```
