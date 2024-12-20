package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/clbanning/mxj/v2"
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
		fmt.Println("outpath", outPath)
		fileName := filepath.Base(path)
		fmt.Println("file", fileName)
		ext := filepath.Ext(fileName)
		outPath = filepath.Join(outPath, fileName[:len(fileName)-len(ext)])

		if action == "xml" && ext == ".toml" {
			outPath += ".xml"

			filesFound++

			r, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("open: %w", err)
			}

			w, err := os.Create(outPath)
			if err != nil {
				return fmt.Errorf("create: %w", err)
			}
			defer w.Close()

			err = tomlToXML(r, w)
			if err != nil {
				return fmt.Errorf("xmlToToml: %w", err)
			}

			return nil
		}

		if action == "toml" && ext == ".xml" {
			outPath += ".toml"

			filesFound++

			fmt.Println("copying", path, "to", outPath)

			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read: %w", err)
			}

			w, err := os.Create(outPath)
			if err != nil {
				return fmt.Errorf("create: %w", err)
			}
			err = xmlToToml(data, w)
			if err != nil {
				return fmt.Errorf("tomlToXML: %w", err)
			}

			return nil
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}

	fmt.Println("XMLs found:", filesFound)

	return nil
}

func xmlToToml(data []byte, w io.Writer) error {
	re := regexp.MustCompile(`(?i)<\?xml[^>]*encoding=["'][^"']*["'][^>]*\?>`)
	data = re.ReplaceAll(data, []byte(`<?xml version="1.0"?>`))

	mxj.CustomDecoder = &xml.Decoder{Strict: false}
	mv, err := mxj.NewMapXml(data)
	if err != nil {
		return fmt.Errorf("new map xml: %w", err)
	}

	enc := toml.NewEncoder(w)

	err = enc.Encode(mv)
	if err != nil {
		return fmt.Errorf("write toml: %w", err)
	}

	return nil
}

func tomlToXML(r io.Reader, w io.Writer) error {
	var v map[string]interface{}
	dec := toml.NewDecoder(r)
	_, err := dec.Decode(&v)
	if err != nil {
		return fmt.Errorf("read toml: %w", err)
	}

	// Convert the map to JSON first
	jsonData, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	// Convert JSON to XML
	mv, err := mxj.NewMapJson(jsonData)
	if err != nil {
		return fmt.Errorf("new map json: %w", err)
	}

	data, err := mv.XmlIndent("", "  ")
	if err != nil {
		return fmt.Errorf("xml indent: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("write xml: %w", err)
	}

	return nil
}

// CharsetReader converts non-UTF-8 encodings to UTF-8
func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	// Handle different charsets (us-ascii is effectively UTF-8 compatible)
	switch strings.ToLower(charset) {
	case "us-ascii", "utf-8":
		return input, nil
	default:
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}
}
