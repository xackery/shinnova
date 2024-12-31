package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xackery/shinnova/nova"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed to run:", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 3 {
		fmt.Println("usage: shinnova <inpath> <outpath>")
		os.Exit(0)
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	fi, err := os.Stat(inPath)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	if !fi.IsDir() {
		return fmt.Errorf("not a directory: %s", inPath)
	}

	fi, err = os.Stat(outPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat: %w", err)
		}
		err = os.MkdirAll(outPath, 0755)
		if err != nil {
			return fmt.Errorf("mkdir: %w", err)
		}
	}

	if !fi.IsDir() {
		return fmt.Errorf("not a directory: %s", outPath)
	}

	xmlFiles, err := filepath.Glob("*.xml")
	if err != nil {
		return fmt.Errorf("glob: %w", err)
	}

	nvFiles, err := filepath.Glob("*.nv")
	if err != nil {
		return fmt.Errorf("glob: %w", err)
	}

	if len(xmlFiles) == 0 && len(nvFiles) == 0 {
		return fmt.Errorf("no files found in source directory")
	}

	action := "xml"
	if len(xmlFiles) == len(nvFiles) {
		return fmt.Errorf("both xml and nv files found, unsure which to convert (delete one set)")
	}

	if len(nvFiles) > len(xmlFiles) {
		action = "nv"
	}

	switch action {
	case "xml":
		fmt.Println("Converting *.xml to *.nv")
		for _, file := range xmlFiles {
			err = convertToNV(filepath.Join(inPath, file), outPath)
			if err != nil {
				return fmt.Errorf("convert %s to nv: %w", file, err)
			}
		}
	case "nv":
		fmt.Println("Converting *.nv to *.xml")
		for _, file := range nvFiles {
			err = convertToXML(filepath.Join(inPath, file), outPath)
			if err != nil {
				return fmt.Errorf("convert %s to xml: %w", file, err)
			}
		}
	}

	return nil
}

func convertToNV(inPath string, outPath string) error {
	fileName := filepath.Base(inPath)
	outPath = filepath.Join(outPath, fileName[:len(fileName)-4]) + ".nv"

	nv, err := nova.NewFromXML(inPath)
	if err != nil {
		return fmt.Errorf("nova parse %s: %w", fileName, err)
	}

	err = nv.ToNV(outPath)
	if err != nil {
		return fmt.Errorf("to nv: %w", err)
	}

	return nil
}

func convertToXML(inPath string, outPath string) error {
	fileName := filepath.Base(inPath)
	outPath = filepath.Join(outPath, fileName[:len(fileName)-3]) + "xml"

	nv, err := nova.NewFromNV(inPath)
	if err != nil {
		return fmt.Errorf("nova parse %s: %w", fileName, err)
	}

	err = nv.ToXML(outPath)
	if err != nil {
		return fmt.Errorf("to xml: %w", err)
	}

	return nil
}
