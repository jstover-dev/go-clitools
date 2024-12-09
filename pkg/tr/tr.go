package tr

import "fmt"

func Tr(argv []string) {
	fmt.Println("> tr");
	for _, arg := range argv {
		fmt.Println(arg);
	}
}
