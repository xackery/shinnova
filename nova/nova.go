package nova

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
)

// Nova represents the UI
type Nova struct {
	Elements []ElementReadWriter
}

// ElementReadWriter represents an element in the UI
type ElementReadWriter interface {
}

func NewFromXML(path string) (*Nova, error) {
	e := &Nova{}
	err := e.FromXML(path)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Nova) FromXML(path string) error {

	bdata, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	data := string(bdata)

	re := regexp.MustCompile(`(?i)<\?xml[^>]*encoding=["'][^"']*["'][^>]*\?>`)
	data = re.ReplaceAllString(data, "")

	dec := xml.NewDecoder(strings.NewReader(string(data)))

	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("token: %w", err)
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "Button":
				o := Button{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 button attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected button attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode button: %w", err)
				}
				e.Elements = append(e.Elements, o)
			case "Label":
				o := Label{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 label attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected label attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode label: %w", err)
				}
				e.Elements = append(e.Elements, o)

			case "STMLbox":
				o := STMLBox{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 STMLbox attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected STMLbox attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode stmlbox: %w", err)
				}
				e.Elements = append(e.Elements, o)

			case "Screen":
				o := Screen{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 screen attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected screen attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode screen: %w", err)
				}
				e.Elements = append(e.Elements, o)
			case "Page":
				o := Page{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 page attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected page attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode page: %w", err)
				}
				e.Elements = append(e.Elements, o)
			case "TabBox":
				o := TabBox{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 tabbox attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected tabbox attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode tabbox: %w", err)
				}
				e.Elements = append(e.Elements, o)
			case "Composite":
				o := Composite{}
				if len(se.Attr) != 0 {
					return fmt.Errorf("expected 0 composite attributes, got %d: %+v", len(se.Attr), se.Attr)
				}

				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode composite: %w", err)
				}

				e.Elements = append(e.Elements, o)
			case "Listbox":
				o := ListBox{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 listbox attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected listbox attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode listbox: %w", err)
				}
				e.Elements = append(e.Elements, o)
			case "Editbox":
				o := EditBox{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 editbox attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected editbox attribute item, got %s", se.Attr[0].Name.Local)
				}
				o.Item = se.Attr[0].Value
				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode editbox: %w", err)
				}
				e.Elements = append(e.Elements, o)

			case "XML":
				continue
			case "Schema":
				continue
			default:
				return fmt.Errorf("unknown element: %s", se.Name.Local)
			}

			//fmt.Printf("element: %#v\n", element)
		case xml.EndElement:

		case xml.CharData:
		}

	}

	//	fmt.Printf("elements: %#v\n", e.Elements)

	return nil
}

func (e *Nova) ToXML(path string) error {
	w, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer w.Close()
	w.WriteString(`<?xml version="1.0" encoding="us-ascii"?>` + "\n")
	w.WriteString(`<XML ID="EQInterfaceDefinitionLanguage">` + "\n")
	tab := 1
	w.WriteString(strings.Repeat("\t", tab) + `<Schema xmlns="EverQuestData" xmlns:dt="EverQuestDataTypes" />` + "\n")

	for _, element := range e.Elements {
		switch e := element.(type) {
		case Button:
			output, err := xml.MarshalIndent(e, strings.Repeat("\t", tab), "\t")
			if err != nil {
				return fmt.Errorf("marshal button: %w", err)
			}
			w.Write(output)
		}
	}

	w.WriteString(`</XML>` + "\n")
	return nil
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

			// if it's an array, repeat break it down
			if v.Field(i).Kind() == reflect.Slice {
				for j := 0; j < v.Field(i).Len(); j++ {
					key := fmt.Sprintf("%s%d", t.Field(i).Name, j)
					value := v.Field(i).Index(j).Interface()
					w.WriteString(fmt.Sprintf("%s%s %v\n", strings.Repeat("\t", tab), key, value))
				}
				continue
			}
			w.WriteString(fmt.Sprintf("%s%s %v\n", strings.Repeat("\t", tab), key, value))
		}
		w.WriteString("\n")
	}
	return nil
}
