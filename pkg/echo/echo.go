package echo

import "fmt"

func Echo(argv []string) {

	printArgs := argv[0:]

	if len(printArgs) > 0 {
		fmt.Print(argv[0])
		for _, arg := range argv[1:] {
			fmt.Printf(" %s", arg)
		}
		fmt.Println();
	}
}


