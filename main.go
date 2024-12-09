package main

import (
	"fmt"
	"os"
	"path"

	"codeberg.org/jstover/clitools/cat"
	"codeberg.org/jstover/clitools/tr"
)

var applications = map[string]func([]string){
	"cat": cat.Cat,
	"tr": tr.Tr,
}

func main() {
	fmt.Println("wat")
	fmt.Println(os.Args[0])

	// Determine program from argv[0] or argv[1]
	for i := range(2) {
		if len(os.Args) > i {
			fn, ok := applications[path.Base(os.Args[i])]
			if ok {
				fn(os.Args[i + 1:])
			}
		}
	}

}
