package main

import (
	"io"
	"json_validator/json_parser"
	"json_validator/printers"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printers.Error("usage: ./json_validator <filename>")
		os.Exit(1);
	}

	filename := os.Args[1]
	printers.Success("filename: " + filename)

	file, err := os.Open(filename)
	if err != nil {
		printers.Error("cannot open file " + filename)
		os.Exit(1);
	}
	printers.Success("file " + filename + " opened!")

	content, err := io.ReadAll(file)
	if err != nil {
		printers.Error("cannot read from file " + filename)
		os.Exit(1);
	}
	printers.Success("content of " + filename + " retrieved!")

	res := json_parser.ParseJSON(string(content))
	if res {
		printers.Success("content of " + filename + " is valid json!")
	} else {
		printers.Success("content of " + filename + " is not valid json!")
	}

	os.Exit(0);
}
