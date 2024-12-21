package nova

import (
	"testing"
)

func TestNovaXML(t *testing.T) {
	n, err := NewFromXML("../default/EQLSUI.xml")
	if err != nil {
		t.Fatalf("NewFromXML: %v", err)
	}
	//fmt.Printf("nova: %#v\n", n)
	err = n.ToNV("../out.nv")
	if err != nil {
		t.Fatalf("ToNV: %v", err)
	}

}
