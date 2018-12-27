package iostuff

import (
	"io"
)

type readwriter struct {
	r io.Reader
	w io.Writer
}

func (self readwriter) Read(p []byte) (int, error) {
	return self.r.Read(p)
}

func (self readwriter) Write(p []byte) (int, error) {
	return self.w.Write(p)
}

func Combine(w io.Writer, r io.Reader) io.ReadWriter {
	return readwriter{
		r: r,
		w: w,
	}
}
