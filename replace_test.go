package replace

import (
	"bytes"
	"log"
	"testing"
)

func TestWriter(t *testing.T) {
	tests := []struct {
		from, to string
		in, want string
	}{
		{
			from: "foo",
			to:   "bar",
			in:   "foo",
			want: "bar",
		},
		{
			from: "foo",
			to:   "bar",
			in:   "foox",
			want: "barx",
		},
		{
			from: "foo",
			to:   "bar",
			in:   "xfooy",
			want: "xbary",
		},
		{
			from: "foo",
			to:   "bar",
			in:   "xfoo",
			want: "xbar",
		},
		{
			from: "fo\no",
			to:   "b\nar",
			in:   "xfo\noy",
			want: "xb\nary",
		},
		{
			from: "fo\no",
			to:   "b\nar",
			in:   "xfoy",
			want: "xfoy",
		},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		w := NewWriter(&buf, test.from, test.to)
		_, err := w.Write([]byte(test.in))
		if err != nil {
			log.Fatal(err)
		}
		w.Close()
		got := buf.String()
		if got != test.want {
			t.Fatalf("want %q for %q but got %q", test.want, test.in, got)
		}
	}
}
