Goで WriteRune を持たない io.Writer で最速で1文字出力するには
==============================================================

WriteRune を持たない io.Writer 相手に 1-rune を出力するのに、fmt.Fprintf(〜,"%c") を使いがちだが、utf8.EncodeRune を使ってやった方が速いのではないかと思い検証してみた。

一応、参考データとして、次のように bufio.NewWriter も含めた形でベンチマークをとってみる。

1. fmt.Fprintf
2. utf8.EncodeRune
3. bufio.NewWriter (インスタンスを使い捨て)
4. bufio.NewWriter (インスタンス使いまわし)

テストプログラム
----------------

`main.go`

```main.go
package writerune

import (
    "bufio"
    "fmt"
    "io"
    "unicode/utf8"
)

func WriteRune1(c rune, w io.Writer) {
    fmt.Fprintf(w, "%c", c)
}

func WriteRune2(c rune, w io.Writer) {
    var buffer [utf8.UTFMax]byte
    n := utf8.EncodeRune(buffer[:], c)
    w.Write(buffer[:n])
}

func WriteRune3(c rune, w io.Writer) {
    br := bufio.NewWriter(w)
    br.WriteRune(c)
    br.Flush()
}

func WriteRune4(c rune, bw interface{ WriteRune(rune) (int, error) }) {
    bw.WriteRune(c)
}
```

`main_test.go`

```main_test.go
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
```

ベンチマーク結果
----------------

`go test -bench . -benchmem`

```go test -bench . -benchmem|
goos: windows
goarch: amd64
pkg: github.com/hymkor/study-go-write-rune
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkWriteRune1-4   	15376118	        80.20 ns/op	       4 B/op	       1 allocs/op
BenchmarkWriteRune2-4   	51887179	        21.47 ns/op	       4 B/op	       1 allocs/op
BenchmarkWriteRune3-4   	  900706	      1322 ns/op	    4096 B/op	       1 allocs/op
BenchmarkWriteRune4-4   	162083506	         7.421 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/hymkor/study-go-write-rune	6.576s
```

結論
----

ほぼ、予想どおり

bufio.NewWriter(使い回し) ＞ utf8.EncodeRune ＞ fmt.Fprintf ＞＞ bufio.NewWriter(使い捨て)

いちいち、utf8.EncodeRune を使う価値はある
