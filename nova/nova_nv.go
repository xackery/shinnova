package nova

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/xackery/shinnova/nxml"
)

var regexLine = regexp.MustCompile(`"([^"]*)"|(\S+)`)

type novaReadToken struct {
	basePath       string
	lineNumber     int
	buf            *bytes.Buffer
	totalLineCount int // will be higher than lineNumber due to includes
}

func NewFromNV(path string) (*Nova, error) {
	e := &Nova{}
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	a := &novaReadToken{
		buf: bytes.NewBuffer(buf),
	}
	err = a.readElements()
	if err != nil {
		return nil, fmt.Errorf("%s:%d %w", path, a.lineNumber, err)
	}

	return e, nil
}

func (e *Nova) ToNV(path string) error {
	w, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer w.Close()

	for _, element := range e.Elements {
		v := reflect.ValueOf(element)
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			tab := 1

			key := t.Field(i).Name
			value := v.Field(i).Interface()
			if i == 0 {
				tab = 0
				key = t.Name()
			}

			switch v.Field(i).Kind() {
			case reflect.Slice:
				key = t.Field(i).Name
				w.WriteString(fmt.Sprintf("%sNum%s %d\n", strings.Repeat("\t", tab), key, v.Field(i).Len()))
				tab = 2
				for j := 0; j < v.Field(i).Len(); j++ {
					value := v.Field(i).Index(j).Interface()
					w.WriteString(fmt.Sprintf("%s%s %v\n", strings.Repeat("\t", tab), key, value))
				}
				tab = 1
				continue
			case reflect.String:
				w.WriteString(fmt.Sprintf("%s%s \"%v\"\n", strings.Repeat("\t", tab), key, value))
				continue
			case reflect.Array:
				switch val := value.(type) {
				case [3]*int:
					w.WriteString(fmt.Sprintf("%s%s ", strings.Repeat("\t", tab), key))
					for i := 0; i < 3; i++ {
						if val[i] != nil {
							w.WriteString(fmt.Sprintf("%v", *val[i]))
						} else {
							w.WriteString("<nil>")
						}
						if i < 2 {
							w.WriteString(" ")
						} else {
							w.WriteString("\n")
						}
					}

					continue
				case [2]*int:
					w.WriteString(fmt.Sprintf("%s%s ", strings.Repeat("\t", tab), key))
					for i := 0; i < 2; i++ {
						if val[i] != nil {
							w.WriteString(fmt.Sprintf("%v", *val[i]))
						} else {
							w.WriteString("<nil>")
						}
						if i < 1 {
							w.WriteString(" ")
						} else {
							w.WriteString("\n")
						}
					}
					continue
				default:
					return fmt.Errorf("unknown array type: %T", value)
				}

			case reflect.Ptr:
				switch val := value.(type) {
				case *string:
					if val == nil {
						w.WriteString(fmt.Sprintf("%s%s <nil>\n", strings.Repeat("\t", tab), key))
						continue
					}
					w.WriteString(fmt.Sprintf("%s%s \"%v\"\n", strings.Repeat("\t", tab), key, *val))
					continue
				case *int:
					if val == nil {
						w.WriteString(fmt.Sprintf("%s%s <nil>\n", strings.Repeat("\t", tab), key))
						continue
					}
					w.WriteString(fmt.Sprintf("%s%s %v\n", strings.Repeat("\t", tab), key, *val))
					continue

				default:
					return fmt.Errorf("unsupported pointer type: %T", value)
				}
			default:
				return fmt.Errorf("unknown kind: %s", v.Field(i).Kind())
			}

		}
		w.WriteString("\n")
	}
	return nil
}

func (a *novaReadToken) readElements() error {
	element := ""
	for {
		args, err := a.readSegmentedLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("read: %w", err)
		}
		if len(args) == 0 {
			continue
		}

		element = strings.ToUpper(string(args[0]))

		if strings.HasPrefix(element, "//") {
			continue
		}

		if strings.HasPrefix(element, "Button") {
			e := &nxml.Button{}
			err = a.readElement(e, args[1])
			if err != nil {
				return fmt.Errorf("%s:%d %w", element, a.lineNumber, err)
			}

			continue
		}

		return fmt.Errorf("unknown element: %s", element)

	}

	if element != "" {
		return fmt.Errorf("unknown element: %s", element)
	}
	return nil
}

func (a *novaReadToken) readSegmentedLine() ([]string, error) {
	line, err := a.readLine()
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
		if len(line) == 0 {
			return nil, err
		}
	}
	matches := regexLine.FindAllStringSubmatch(line, -1)
	args := []string{}
	for _, match := range matches {
		if match[2] == "//" {
			break
		}
		if match[1] != "" {
			args = append(args, match[1])
		} else {
			args = append(args, match[2])
		}
	}
	return args, nil
}

// Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
func (a *novaReadToken) readLine() (string, error) {
	line := ""
	p := make([]byte, 1)
	for {
		_, err := a.buf.Read(p)
		if err != nil {
			if err == io.EOF {
				a.lineNumber++
				return line, err
			}

			return "", err
		}
		if p[0] != '\n' {
			line += string(p)
			continue
		}
		a.lineNumber++
		if strings.HasPrefix(strings.TrimSpace(line), "//") {
			line = ""
			continue
		}
		if strings.TrimSpace(line) == "" {
			line = ""
			continue
		}
		return line, nil
	}
}

func (a *novaReadToken) readElement(e interface{}, element string) error {
	v := reflect.ValueOf(e).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Addr().Interface()
		err := a.readField(value, key)
		if err != nil {
			return fmt.Errorf("read field: %w", err)
		}
	}

	return nil
}

func (a *novaReadToken) readField(value interface{}, key string) error {
	args, err := a.readSegmentedLine()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	if len(args) == 0 {
		return fmt.Errorf("expected %s, got nothing", key)
	}
	if !strings.EqualFold(args[0], key) {
		return fmt.Errorf("expected %s, got %s", key, args[0])
	}
	if len(args) < 2 {
		return fmt.Errorf("expected %s value, got nothing", key)
	}
	err = setField(value, args[1])
	if err != nil {
		return fmt.Errorf("set field: %w", err)
	}
	return nil
}

func setField(value interface{}, arg string) error {
	v := reflect.ValueOf(value).Elem()
	switch v.Kind() {
	case reflect.String:
		v.SetString(arg)
	case reflect.Int:
		i, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return fmt.Errorf("parse int: %w", err)
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(arg)
		if err != nil {
			return fmt.Errorf("parse bool: %w", err)
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported type: %s", v.Kind())
	}
	return nil
}
