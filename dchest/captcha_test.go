package dchest

import "testing"

func TestGenerate(t *testing.T) {
	m := Generate()
	if m["code"] != 1 {
		t.Fatalf("not expected result to be code %d", m["code"])
	}
}
