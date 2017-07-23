package main

import (
	"fmt"
	"os"
)

// printHelp prints the usage help to standard error and terminates the process.
func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "Compress:    %s [c] [0|1|2|3] < <file>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Decompress:  %s [d] [0|1|2|3] < <file>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Analyze:     %s [a] < <file>\n", os.Args[0])
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printHelp()
	}
	switch os.Args[1] {
	case "a":
		subargs := args[1:]
		if len(subargs) != 0 {
			printHelp()
		}
		analyze()
	case "c":
		subargs := args[1:]
		if len(subargs) != 1 {
			printHelp()
		}
		switch subargs[0] {
		case "0":
			compress0()
		case "1":
			compress1()
		case "2":
			compress2()
		case "3":
			compress3()
		default:
			printHelp()
		}
	case "d":
		subargs := args[1:]
		if len(subargs) != 1 {
			printHelp()
		}
		switch subargs[0] {
		case "0":
			decompress0()
		case "1":
			decompress1()
		case "2", "3":
			fmt.Fprintf(os.Stderr, "Decompression unsupported with format #%s", os.Args[2])
		default:
			printHelp()
		}
	default:
		printHelp()
	}
}
