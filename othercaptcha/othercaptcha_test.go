package othercaptcha

import "testing"

func BenchmarkGenerateV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := Generate()
		if m["code"] != 1 {
			b.Fatalf("not expected result to be code %d", m["code"])
		}
	}
}
