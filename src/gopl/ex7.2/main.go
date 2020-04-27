package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf bytes.Buffer
	wr, count := CountingWriter(&buf)
	fmt.Fprintf(wr, "Hello %s", "world\n")
	fmt.Println(*count)
	wr.Write([]byte("Another string"))
	fmt.Println(*count)
}

type newWriter struct {
	counter int64
	writer  io.Writer
}

func (nw *newWriter) Write(p []byte) (int, error) {
	written, err := nw.writer.Write(p)
	nw.counter += int64(written)
	return written, err
}

func newNewWriter(w io.Writer) *newWriter {
	return &newWriter{
		counter: 0,
		writer:  w,
	}
}

// CountingWriter see p 174
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var nw = newNewWriter(w)
	return nw, &(nw.counter)
}
