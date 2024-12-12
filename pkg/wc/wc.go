package wc

import (
	"fmt"
	"io"
	"os"
)

const BUFFER_SIZE = 4096

type CountResult struct {
	filename string
	newlines int
	words    int
	bytes    int
}

func Wc(argv []string) {

	for _, arg := range argv {

		result := CountResult{
			filename: arg,
			newlines: 0,
			words:    1,
			bytes:    0,
		}

		file, err := os.Open(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		buf := make([]byte, BUFFER_SIZE)
		for {
			n, err := file.Read(buf)
			if err == io.EOF {
				break
			}

			result.bytes += n

			for _, char := range buf {
				if char == '\n' {
					result.newlines++
				} else if char == ' ' {
					result.words++
				}

			}

		}
		fmt.Printf("%d %d %d %s\n", result.newlines, result.words, result.bytes, result.filename)
	}

}
