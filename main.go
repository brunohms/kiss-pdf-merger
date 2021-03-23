package main

import (
	"fmt"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var inputPaths []string
var outputPath string
var revertNames map[string]string

func changeTextToNumbers(text string) string {
	text = strings.Replace(text, "um", "1", -1)
	text = strings.Replace(text, "dois", "2", -1)
	text = strings.Replace(text, "três", "3", -1)
	text = strings.Replace(text, "tres", "3", -1)
	text = strings.Replace(text, "quatro", "4", -1)
	text = strings.Replace(text, "cinco", "5", -1)
	text = strings.Replace(text, "seis", "6", -1)
	text = strings.Replace(text, "sete", "7", -1)
	text = strings.Replace(text, "oito", "8", -1)
	text = strings.Replace(text, "nove", "9", -1)
	text = strings.Replace(text, "dez", "10", -1)
	text = strings.Replace(text, "onze", "11", -1)
	text = strings.Replace(text, "doze", "12", -1)
	text = strings.Replace(text, "treze", "13", -1)
	text = strings.Replace(text, "quatorze", "14", -1)
	text = strings.Replace(text, "catorze", "14", -1)
	text = strings.Replace(text, "quinze", "15", -1)
	text = strings.Replace(text, "dezesseis", "16", -1)
	text = strings.Replace(text, "dezessete", "17", -1)
	text = strings.Replace(text, "dezoito", "18", -1)
	text = strings.Replace(text, "dezenove", "19", -1)
	text = strings.Replace(text, "vinte", "20", -1)
	text = strings.Replace(text, "trinta", "30", -1)
	text = strings.Replace(text, "quarenta", "40", -1)
	text = strings.Replace(text, "cinquenta", "50", -1)
	text = strings.Replace(text, "sessenta", "60", -1)
	text = strings.Replace(text, "setenta", "70", -1)
	text = strings.Replace(text, "oitenta", "80", -1)
	text = strings.Replace(text, "noventa", "90", -1)

	return text
}

func walkFunc(path string, info os.FileInfo, _ error) error {
	if strings.HasSuffix(path, ".pdf") && info.Size() > 0 && path != outputPath {
		absolutePath, _ := filepath.Abs(path)
		originalPathName := absolutePath
		absolutePath = changeTextToNumbers(absolutePath)
		revertNames[absolutePath] = originalPathName
		inputPaths = append(inputPaths, absolutePath)
	}
	return nil
}

func main() {
	path := "./proc/"

	outputPath, _ = filepath.Abs(path)
	outputPath = filepath.Base(outputPath) + ".pdf"

	revertNames = make(map[string]string)

	err := filepath.Walk(path, walkFunc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(inputPaths) == 0 {
		fmt.Println("No pdf files found.")
		return
	}

	sort.Strings(inputPaths)
	for i, path := range inputPaths {
		fmt.Println("Path:", path)
		inputPaths[i] = revertNames[path]
	}

	err = pdf.MergeCreateFile(inputPaths, outputPath, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Finished merging files, result at", outputPath)
}