package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./json_validator <filepath>");
		os.Exit(1);
	}

	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Cannot open file: " + filepath);
		os.Exit(1);
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Cannot read file: " + filepath);
		os.Exit(1)
	}
	fmt.Printf("%s\n", content)
}
