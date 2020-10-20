package main

import (
	"fmt"

	"github.com/mingderwang/hello/morestrings"
        "github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
        fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
