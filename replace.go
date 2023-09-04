package replace

import (
	"bytes"
	"io"

	"golang.org/x/text/transform"
)

type replacer struct {
	transform.NopResetter
	from []byte
	to   []byte
}

var _ transform.Transformer = (*replacer)(nil)

func (t *replacer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	if len(src) == 0 && atEOF {
		return
	}
	lTo := len(t.to)
	lFrom := len(t.from)

	for nDst < len(dst) && nSrc < len(src) {
		if len(src) >= lFrom && bytes.HasPrefix(src[nSrc:], t.from) {
			if nDst+lTo > len(dst) {
				err = transform.ErrShortDst
				break
			}
			n := copy(dst[nDst:], t.to)
			if n <= 0 {
				break
			}
			if n < lFrom || lFrom < lTo {
				nSrc += lFrom
			} else {
				nSrc += n
			}
			nDst += n
		} else {
			dst[nDst] = src[nSrc]
			nSrc++
			nDst++
		}
	}
	return
}

func NewWriter(w io.Writer, from, to string) io.WriteCloser {
	return transform.NewWriter(w, &replacer{
		from: []byte(from),
		to:   []byte(to),
	})
}

func NewReader(r io.Reader, from, to string) io.Reader {
	return transform.NewReader(r, &replacer{
		from: []byte(from),
		to:   []byte(to),
	})
}
