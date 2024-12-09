package cat

import "fmt"

func Cat(argv []string) {
	for _, arg := range argv {
		fmt.Println(arg)
	}
}
