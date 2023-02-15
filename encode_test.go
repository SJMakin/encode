package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input string
}{
	{input: "| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$"},
	{input: "| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$"},
	{input: `| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$
    | || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$`},
}

func TestEncodeDecode(t *testing.T) {
	in := "| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$"
	encoded := EncodeOptimized(in)
	out := DecodeOptimized(encoded)
	if in != out {
		t.Error("No match.")
	}
}

func benchEncode(b *testing.B, f func(string) string) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(v.input)
			}
		})
	}
}

func benchDecode(b *testing.B, f func(string) string) {
	for _, v := range table {
		encoded := EncodeOptimized(v.input)
		b.Run(fmt.Sprintf("input_size_%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(encoded)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	benchEncode(b, Encode)
}

func BenchmarkEncodeOptimized(b *testing.B) {
	benchEncode(b, EncodeOptimized)
}

func BenchmarkDecode(b *testing.B) {
	benchDecode(b, Decode)
}

func BenchmarkDecodeOptimized(b *testing.B) {
	benchDecode(b, DecodeOptimized)
}

func BenchmarkDecodeOptimizedSimplified(b *testing.B) {
	benchDecode(b, DecodeOptimizedSimplified)
}
