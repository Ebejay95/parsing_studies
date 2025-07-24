package main

import (
	"io"
	"parsing_studies/json_parser"
	"parsing_studies/printers"
	"os"
)

func applyParsingMode(mode string, filename string, content string) {
	var res bool = false
	switch mode {
		case "json":
			res = json_parser.IsValidJSON(content)
		default:
			printers.Error(mode + " mode is not supported!")
			return
	}
	if res {
		printers.Success("content of " + filename + " is valid " + mode + "!")
	} else {
		printers.Error("content of " + filename + " is not valid " + mode + "!")
	}
}

func main() {
	if len(os.Args) != 3 {
		printers.Error("usage: ./parsing_studies <mode> <filename>")
		os.Exit(1)
	}

	mode := os.Args[1]
	filename := os.Args[2]
	printers.Log("mode: " + mode + ", filename: " + filename)

	file, err := os.Open(filename)
	if err != nil {
		printers.Error("cannot open file " + filename)
		os.Exit(1)
	}
	printers.Log("file " + filename + " opened!")

	content, err := io.ReadAll(file)
	if err != nil {
		printers.Error("cannot read from file " + filename)
		os.Exit(1)
	}
	printers.Log("content of " + filename + " retrieved!")

	applyParsingMode(mode, filename, string(content))

	os.Exit(0)
}