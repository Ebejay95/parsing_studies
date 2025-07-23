package main

import "fmt"
import "os"
import "io"
import "strings"
import "encoding/json"

func isValidJSON(json string) bool {
	var js json.RawMessage
	return json.Unma
}

func main(){
	if len(os.Args) != 2 {
		fmt.Println("usage: ./app <filename>")
		os.Exit(1)
	}

	filename := os.Args[1];

	fmt.Println("filename is: " + filename)

	if !strings.HasSuffix(filename, ".json") {
		fmt.Println("usage: with json files only")
		os.Exit(1)
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("could not open file: " + filename)
		os.Exit(1)
	}

	content, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("could not read file: " + filename)
		os.Exit(1)
	}

	fmt.Printf("%s\n", content)

	os.Exit(0)
}