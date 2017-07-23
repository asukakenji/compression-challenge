package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage: go run main.go [c|d] [0|1|2|3] < file")
	fmt.Println("c for compress")
	fmt.Println("d for decompress")
}

func main() {
	if len(os.Args) < 3 {
		printHelp()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "c":
		switch os.Args[2] {
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
		switch os.Args[2] {
		case "0", "1", "2", "3":
			fmt.Fprintf(os.Stderr, "Decompression unsupported with format #%s", os.Args[2])
		default:
			printHelp()
		}
	default:
		printHelp()
	}
}
