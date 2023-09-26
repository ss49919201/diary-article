package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/imports"
)

//go:generate go run main.go $FILE_NAME

func main() {
	var (
		fileName      = os.Args[1]
		interfaceName string
		structName    string
	)

	for i, v := range strings.Split(fileName, "_") {
		title := cases.Title(language.Armenian).String(v)
		interfaceName += title
		if i == 0 {
			structName += v
		} else {
			structName += title
		}
	}

	tmpl, err := template.ParseFiles("handler.gotmpl")
	if err != nil {
		log.Fatal("Failed to parse template: ", err)
	}

	buf := bytes.NewBuffer(nil)

	if err := tmpl.Execute(buf, map[string]any{
		"InterfaceName": interfaceName,
		"StructName":    structName,
	}); err != nil {
		log.Fatal("Failed to execute template: ", err)
	}

	bytes, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		log.Fatal("Failed to format and adjust imports source: ", err)
	}

	var path string
	_, filename, _, _ := runtime.Caller(0)
	for dir := filepath.Dir(filename); dir != "/" && dir != "."; dir = filepath.Dir(dir) {
		if filepath.Base(dir) == "diary" {
			path = filepath.Join(dir, "internal", "handler", fileName+".go")
		}
	}

	if path == "" {
		log.Fatal("target directory is not found")
	}

	if err := os.WriteFile(path, bytes, 0644); err != nil {
		log.Fatal("Failed to write file: ", err)
	}
}
