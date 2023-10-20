package writerune

import (
	"bufio"
	"io"
	"testing"
)

func BenchmarkWriteRune1(B *testing.B) {
	for i := 0; i < B.N; i++ {
		WriteRune1('あ', io.Discard)
	}
}

func BenchmarkWriteRune2(B *testing.B) {
	for i := 0; i < B.N; i++ {
		WriteRune2('あ', io.Discard)
	}
}

func BenchmarkWriteRune3(B *testing.B) {
	for i := 0; i < B.N; i++ {
		WriteRune3('あ', io.Discard)
	}
}

func BenchmarkWriteRune4(B *testing.B) {
	bw := bufio.NewWriter(io.Discard)
	for i := 0; i < B.N; i++ {
		WriteRune4('あ', bw)
	}
}
