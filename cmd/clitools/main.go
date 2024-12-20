package main

import (
	"os"
	"path"

	"codeberg.org/jstover/clitools/pkg/cat"
	"codeberg.org/jstover/clitools/pkg/echo"
	"codeberg.org/jstover/clitools/pkg/tr"
	"codeberg.org/jstover/clitools/pkg/wc"
)

var applications = map[string]func([]string){
	"cat": cat.Cat,
	"echo": echo.Echo,
	"tr": tr.Tr,
	"wc": wc.Wc,
}

func main() {
	for i := range(2) {
		if len(os.Args) > i {
			fn, ok := applications[path.Base(os.Args[i])]
			if ok {
				os.Args = os.Args[1:]
				fn(os.Args)
			}
		}
	}
}
