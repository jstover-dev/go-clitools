package main

import (
	"os"

	"codeberg.org/jstover/clitools/pkg/cat"
)

func main(){
	cat.Cat(os.Args)
}
