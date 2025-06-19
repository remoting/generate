package pkg

import (
	"log"
	"os"
	"path/filepath"
)

func Envs() (string, string, string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	pkgName := os.Getenv("GOPACKAGE")
	fileName := os.Getenv("GOFILE")

	if pkgName == "" {
		log.Fatal("Error: GOPACKAGE environment variable not set. Are you running this with go generate?")
	}
	if fileName == "" {
		log.Fatal("Error: GOFILE environment variable not set. Are you running this with go generate?")
	}
	return filepath.Join(cwd, fileName), pkgName, fileName
}

func GenExecute(gen Generator) {
	sourceAbsPath, _package, _file := Envs()
	outputFileName := gen.Name(_package, _file)
	outputFilePath := filepath.Join(filepath.Dir(sourceAbsPath), outputFileName)
	f, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Error creating file %s: %v", outputFilePath, err)
	}
	defer f.Close()

	_content := gen.Gen(_package, _file)

	_, err = f.WriteString(_content)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", outputFilePath, err)
	}
}
