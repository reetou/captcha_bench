package bench

import (
	"captcha/captcha"
	"captcha/dchest"
	"captcha/lifei6671"
	"captcha/othercaptcha"
	"testing"
)

func BenchmarkGenerate_mojocn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := captcha.Generate()
		if m["code"] != 1 {
			b.Fatalf("not expected result to be code %d", m["code"])
		}
	}
}

func BenchmarkGenerate_steambap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := othercaptcha.Generate()
		if m["code"] != 1 {
			b.Fatalf("not expected result to be code %d", m["code"])
		}
	}
}

func BenchmarkGenerate_lifei6671(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := lifei6671.Generate()
		if m["code"] != 1 {
			b.Fatalf("not expected result to be code %d", m["code"])
		}
	}
}

func BenchmarkGenerate_dchest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := dchest.Generate()
		if m["code"] != 1 {
			b.Fatalf("not expected result to be code %d", m["code"])
		}
	}
}
