package main

import (
	"fmt"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	"os"
	"path/filepath"
	"strings"
)

var inputPaths []string
var outputPath string

func walkFunc(path string, info os.FileInfo, _ error) error {
	if strings.HasSuffix(path, ".pdf") && info.Size() > 0 && !strings.HasPrefix(path, outputPath){
		absolutePath, _ := filepath.Abs(path)
		inputPaths = append(inputPaths, absolutePath)
	}
	return nil
}

func main() {
	outputPath, _ = filepath.Abs("./")
	outputPath = filepath.Base(outputPath)

	err := filepath.Walk("./", walkFunc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(inputPaths) == 0 {
		fmt.Println("No pdf files found.")
		return
	}

	for _, path := range inputPaths {
		fmt.Println("Path:", path)
	}

	err = pdf.MergeCreateFile(inputPaths, outputPath + ".pdf", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Finished merging files, result at", outputPath)
}