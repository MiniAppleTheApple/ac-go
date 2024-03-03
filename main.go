package main

import (
	"flag"
	"go/parser"
	"go/token"
	"log"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var (
		typeName string
		subtypeName string
		fileName string
		out string = ""
	)

	flag.StringVar(&typeName, "type", "", "Name of the type you wanna generate")
	flag.StringVar(&fileName, "filename", "", "Name of the file that contains the type you are saying")
	flag.StringVar(&subtypeName, "subtype", "", "Name of the subtype you wanna generate")
	flag.StringVar(&out, "out", "", "Name of the subtype you wanna generate")

	flag.Parse()

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, nil, parser.AllErrors)	

	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(file.Decls)
}
