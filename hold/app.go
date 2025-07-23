package main

import "fmt"
import "os"
import "io"
import "strings"
import "encoding/json"
import "json_validator/json_parser"

func isValidJSON(jsonStr string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(jsonStr), &js) == nil
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

	//fmt.Printf("%s\n", content)

    //fmt.Println("Valid JSON (Marshal):", isValidJSON(string(content)))
    fmt.Println("Valid JSON:", json_parser.IsValidJSON(string(content)))
	os.Exit(0)
}