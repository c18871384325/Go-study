package strutil

import (
	"testing"
)

func BenchmarkString2Intv1(b *testing.B) {
	for i := 0; i < 100000; i++ {
		String2Intv1(i)
	}
}

func BenchmarkString2Intv2(b *testing.B) {
	for i := 0; i < 100000; i++ {
		String2Intv2(i)
	}
}

func BenchmarkString2Intv3(b *testing.B) {
	var i int64
	for i = 0; i < 100000; i++ {
		String2Intv3(i)
	}
}
