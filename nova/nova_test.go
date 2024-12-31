package nova

import (
	"fmt"
	"testing"
)

func TestConvertXML(t *testing.T) {
	fileName := "EQUI_PlayerWindow" //"EQUI_CharacterCreate" //"EQUI_BankWnd"
	n, err := NewFromXML(fmt.Sprintf("../default/%s.xml", fileName))
	if err != nil {
		t.Fatalf("NewFromXML: %v", err)
	}
	//fmt.Printf("nova: %#v\n", n)
	err = n.ToNV(fmt.Sprintf("../bin/%s.nv", fileName))
	if err != nil {
		t.Fatalf("ToNV: %v", err)
	}
}

func TestConverNV(t *testing.T) {
	fileName := "EQUI_BankWnd"
	n, err := NewFromXML(fmt.Sprintf("../bin/%s.nv", fileName))
	if err != nil {
		t.Fatalf("NewFromNV: %v", err)
	}
	//fmt.Printf("nova: %#v\n", n)
	err = n.ToXML(fmt.Sprintf("../bin/%s.xml", fileName))
	if err != nil {
		t.Fatalf("ToXML: %v", err)
	}
}
