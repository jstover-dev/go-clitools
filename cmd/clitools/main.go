package main

import (
	"os"
	"path"

	"codeberg.org/jstover/clitools/pkg/cat"
	"codeberg.org/jstover/clitools/pkg/echo"
	"codeberg.org/jstover/clitools/pkg/tr"
)

var applications = map[string]func([]string){
	"cat": cat.Cat,
	"echo": echo.Echo,
	"tr": tr.Tr,
}

func main() {
	for i := range(2) {
		if len(os.Args) > i {
			fn, ok := applications[path.Base(os.Args[i])]
			if ok {
				fn(os.Args[i + 1:])
			}
		}
	}
}
