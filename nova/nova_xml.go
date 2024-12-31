package nova

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/xackery/shinnova/nv"
	"github.com/xackery/shinnova/nxml"
)

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
				src := nxml.Button{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 button attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected button attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode button: %w", err)
				}
				dst, err := nv.NewButton(&src)
				if err != nil {
					return fmt.Errorf("new button %s: %w", se.Attr[0].Name.Local, err)
				}
				e.Elements = append(e.Elements, dst)
			case "Label":
				src := nxml.Label{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 label attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected label attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode label: %w", err)
				}

				dst, err := nv.NewLabel(&src)
				if err != nil {
					return fmt.Errorf("new label %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "InvSlot":
				src := nxml.InvSlot{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 invslot attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected invslot attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode invslot: %w", err)
				}

				dst, err := nv.NewInvSlot(&src)
				if err != nil {
					return fmt.Errorf("new invslot %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "STMLbox":
				o := nxml.STMLBox{}
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
				src := nxml.Screen{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 screen attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected screen attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode screen: %w", err)
				}
				dst, err := nv.NewScreen(&src)
				if err != nil {
					return fmt.Errorf("new screen %s: %w", se.Attr[0].Name.Local, err)
				}
				e.Elements = append(e.Elements, dst)
			case "Page":
				o := nxml.Page{}
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
				o := nxml.TabBox{}
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
				o := nxml.Composite{}
				if len(se.Attr) != 0 {
					return fmt.Errorf("expected 0 composite attributes, got %d: %+v", len(se.Attr), se.Attr)
				}

				err = dec.DecodeElement(&o, &se)
				if err != nil {
					return fmt.Errorf("decode composite: %w", err)
				}

				e.Elements = append(e.Elements, o)
			case "Listbox":
				o := nxml.ListBox{}
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
				o := nxml.Editbox{}
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
			case "Ui2DAnimation":
				src := nxml.Ui2DAnimation{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 ui2danimation attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected ui2danimation attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode ui2danimation: %w", err)
				}

				dst, err := nv.NewUi2DAnimation(&src)
				if err != nil {
					return fmt.Errorf("new ui2danimation %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "StaticAnimation":
				src := nxml.StaticAnimation{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 staticanimation attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected staticanimation attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode staticanimation: %w", err)
				}

				dst, err := nv.NewStaticAnimation(&src)
				if err != nil {
					return fmt.Errorf("new staticanimation %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "ButtonDrawTemplate":
				src := nxml.ButtonDrawTemplate{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 buttondrawtemplate attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected buttondrawtemplate attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode buttondrawtemplate: %w", err)
				}

				dst, err := nv.NewButtonDrawTemplate(&src)
				if err != nil {
					return fmt.Errorf("new buttondrawtemplate %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "Combobox":
				src := nxml.Combobox{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 combobox attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected combobox attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value
				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode combobox: %w", err)
				}

				dst, err := nv.NewCombobox(&src)
				if err != nil {
					return fmt.Errorf("new combobox %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "TextureInfo":
				src := nxml.TextureInfo{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 textureinfo attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected textureinfo attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value

				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode textureinfo: %w", err)
				}

				dst, err := nv.NewTextureInfo(&src)
				if err != nil {
					return fmt.Errorf("new textureinfo %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
			case "Gauge":
				src := nxml.Gauge{}
				if len(se.Attr) != 1 {
					return fmt.Errorf("expected 1 gauge attribute, got %d: %+v", len(se.Attr), se.Attr)
				}
				if se.Attr[0].Name.Local != "item" {
					return fmt.Errorf("expected gauge attribute item, got %s", se.Attr[0].Name.Local)
				}
				src.Item = se.Attr[0].Value

				err = dec.DecodeElement(&src, &se)
				if err != nil {
					return fmt.Errorf("decode gauge: %w", err)
				}

				dst, err := nv.NewGauge(&src)
				if err != nil {
					return fmt.Errorf("new gauge %s: %w", se.Attr[0].Name.Local, err)
				}

				e.Elements = append(e.Elements, dst)
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
		case nxml.Button:
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
