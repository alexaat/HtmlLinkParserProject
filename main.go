package main

import (
	"HtmlLinkParseProject/htmlLinkParser"
	"fmt"
	"os"
)

var f1 = "ex1.html"
var f2 = "ex2.html"
var f3 = "ex3.html"
var f4 = "ex4.html"

func main() {

	reader, err := os.Open(f4)
	if err != nil {
		fmt.Println("Cannot open file")
		return
	}

	defer reader.Close()

	fmt.Printf("%+v\n", htmlLinkParser.ParseFileFromUrl(reader))
}
