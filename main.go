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
