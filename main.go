package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)


type TemplateData struct {
	Content string
	Type string
}

func main() {
	var (
		typeName string
		contentFilePath string
		templatePath string	
	)

	flag.StringVar(&typeName, "type", "", "Name of the type you wanna generate")
	flag.StringVar(&templatePath, "template", "", "Path of the template")
	flag.StringVar(&contentFilePath, "file", "", "Path of the file")

	flag.Parse()

	buf, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)	
	}

	contentBuf, err := os.ReadFile(contentFilePath)
	if err != nil {
		log.Fatal(err)	
	}

	content := string(contentBuf)
	
	builder := strings.Builder{}

	templateContent := string(buf)
	templateData := TemplateData {
		Content: string(content),
		Type: typeName,
	}

	template, err := template.New("template").Parse(templateContent)
	if err != nil {
		log.Fatal(err)
	}

	err = template.Execute(&builder, templateData)
	if err != nil {
		log.Fatal(err)
	}

	mode, err := os.Lstat(contentFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(contentFilePath, []byte(builder.String()), mode.Mode().Perm())
	if err != nil {
		log.Fatal(err)
	}

}
