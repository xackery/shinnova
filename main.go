package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		fmt.Println("usage: shinnova xml|toml <directory>")
		os.Exit(0)
	}

	action := os.Args[1]
	if action != "xml" && action != "toml" {
		return fmt.Errorf("unknown action: %s, acceptable ones are xml or toml", action)
	}
	fmt.Println("action", action)

	directory := os.Args[2]

	fi, err := os.Stat(directory)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	if !fi.IsDir() {
		return fmt.Errorf("not a directory: %s", directory)
	}

	filesFound := 0

	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk: %w", err)
		}

		if info.IsDir() {
			if strings.HasSuffix(path, ".git") {
				return filepath.SkipDir
			}
			return nil
		}

		outPath := filepath.Dir(path)
		fileName := filepath.Base(path)
		ext := filepath.Ext(fileName)
		outPath = filepath.Join(outPath, fileName[:len(fileName)-len(ext)])

		if action == "toml" && ext == ".xml" {
			outPath += ".nv"

			filesFound++

			//		fmt.Println("copying", path, "to", outPath)

			nv, err := nova.NewFromXML(path)
			if err != nil {
				return fmt.Errorf("nova parse %s: %w", fileName, err)
			}

			err = nv.ToNV(outPath)
			if err != nil {
				return fmt.Errorf("to nv: %w", err)
			}

		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("XMLs found:", filesFound)

	return nil
}
