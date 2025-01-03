package cat

import (
	"fmt"
	"bufio"
	"os"
)

type Flags struct {
	A       bool
	b       bool
	E       bool
	n       bool
	s       bool
	T       bool
	u       bool
	v       bool
	help    bool
	version bool
}

type Arguments struct {
	flags  Flags
	inputs []string
}

func parseArgs(argv []string) Arguments {

	flags := Flags{}
	inputs := []string{}

	for _, arg := range argv[1:] {
		if arg == "-A" || arg == "--show-all" {
			flags.A = true
		} else if arg == "-b" || arg == "--number-nonblank" {
			flags.b = true
			flags.n = true
		} else if arg == "-e" {
			flags.v = true
			flags.E = true
		} else if arg == "-E" || arg == "--show-ends" {
			flags.E = true
		} else if arg == "-n" || arg == "--number" {
			flags.n = true
		} else if arg == "-s" || arg == "--squeeze-blank" {
			flags.s = true
		} else if arg == "-t" {
			flags.v = true
			flags.T = true
		} else if arg == "-T" || arg == "--show-tabs" {
			flags.T = true
		} else if arg == "-u" {
			// ignored
		} else if arg == "-v" || arg == "--show-nonprinting" {
			flags.v = true
		} else if arg == "--help" {
			flags.help = true
		} else if arg == "--version" {
			flags.version = true
		} else {
			inputs = append(inputs, arg)
		}
	}
	if len(inputs) == 0 {
		inputs = append(inputs, "-")
	}
	return Arguments{flags: flags, inputs: inputs}
}

func Cat(argv []string) {
	args := parseArgs(argv)

	if (args.flags.help) {
		fmt.Println("Heeeelllp")
		os.Exit(0)
	} else if (args.flags.version) {
		fmt.Println("cat 0.1")
		os.Exit(0)
	}

	prev := byte('\n')
	ignoreNewlines := false

	lineNumber := 1
	outBuffer := bufio.NewWriter(os.Stdout)

	for _, filename := range args.inputs {

		var err error
		var file *os.File

		if filename == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()
		}
		reader := bufio.NewReader(file)

		for {
			chr, err := reader.ReadByte()
			if err != nil {
				break
			}
			if (prev == '\n') {
				if (args.flags.s) {
					if (chr == '\n') {
						if (ignoreNewlines) {
							continue
						}
						ignoreNewlines = true
					} else {
						ignoreNewlines = false
					}
				}
				if (args.flags.n) {
					if (!args.flags.b || chr != '\n') {
						outBuffer.WriteString(fmt.Sprintf("%6d\t", lineNumber))
						lineNumber++
					}
				}
			}

			if (chr == '\n') {
				if (args.flags.E) {
					outBuffer.WriteByte('$')
				}
			} else if (chr == '\t') {
				if (args.flags.T) {
					outBuffer.WriteString("^I")
				}
			} else if (args.flags.v) {
				// if isascii(chr)
				if ! (chr < 0x80) {
					outBuffer.WriteString("M-");
					// toascii
					chr &= 0x7F
				}
				// if iscntrl(chr)
				if (chr <= 0x1F || chr == 0x7F) {
					outBuffer.WriteByte('^')
					if (chr == 0x7F) {
						outBuffer.WriteByte('?')
					} else {
						outBuffer.WriteByte(chr | 0x40)
					}
					continue
				}

			}

			outBuffer.WriteByte(chr)
			prev = chr
		}
		outBuffer.Flush()

	}
}
