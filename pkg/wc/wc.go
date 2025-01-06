package wc

import (
	"fmt"
	"io"
	"os"
	"codeberg.org/jstover/clitools/pkg/common"
	"github.com/alexflint/go-arg"
)

const BUFFER_SIZE = 4096

type CountResult struct {
	filename string
	bytes int
	lines int
	words int
	chars int
}

type Arguments struct {
	Bytes bool			`arg:"-c,--bytes" help:"The number of bytes in each input file is written to the standard output."`
	Lines bool  		`arg:"-l,--lines" help:"The number of lines in each input file is written to the standard output."`
	Words bool 			`arg:"-w,--words" help:"The number of words in each input file is written to the standard output."`
	Chars bool			`arg:"-m,--chars" help:"Count characters instead of bytes"`
	Human bool			`arg:"-h,--human" help:"Use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte, Petabyte, and Exabyte in order to reduce the number of digits to four or fewer using powers of 2 for sizes (K=1024, M=1048576, etc.)."`
	Files []string		`arg:"positional"`
}

func parseArgs() Arguments {
	var args Arguments
	arg.MustParse(&args)
	if len(args.Files) == 0 {
		args.Files = append(args.Files, "-")
	}
	// If no options are specified, defualt to -clw
	if !args.Bytes && !args.Lines && !args.Words {
		args.Bytes = true
		args.Lines = true
		args.Words = true
	}
	return args
}

func (Arguments) Version() string {
	return "wc 1.0.0"
}



func Wc(argv []string) {
	os.Args = argv
	args := parseArgs()

	for _, filename := range args.Files {
		reader, err := common.NewFileReader(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer reader.Close()

		result := CountResult{
			filename: filename,
			bytes: 0,
			chars: 0,
			lines: 0,
			words: 1,
		}

		// Keep track of when we move from one line to the next
		isNewLine := false

		for {
			chr, err := reader.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			result.bytes++
			if chr == '\n' {
				result.lines++
				isNewLine = true
			} else if isNewLine || chr == ' ' {
				result.words++
				isNewLine = false
			}
		}

		if args.Lines {
			fmt.Printf(" %d", result.lines)
		}
		if args.Words {
			fmt.Printf(" %d", result.words)
		}
		if args.Bytes {
			fmt.Printf(" %d", result.bytes)
		}
		if result.filename != "-" {
			fmt.Printf(" %s\n", result.filename)
		} else {
			fmt.Println()
		}
	}
}
