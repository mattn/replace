package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mattn/replace"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func main() {
	var err error
	flag.Parse()
	f := os.Stdin
	switch flag.NArg() {
	case 2:
		writer := replace.NewWriter(os.Stdout, flag.Arg(0), flag.Arg(1))
		_, err = io.Copy(writer, f)
		fatalIf(err)
	case 3:
		f, err := os.Open(flag.Arg(2))
		fatalIf(err)

		fname := filepath.Base(flag.Arg(2)) + "-temp-*"
		tmp, err := ioutil.TempFile(filepath.Dir(flag.Arg(2)), fname)
		fatalIf(err)
		defer func() {
			err := os.Rename(tmp.Name(), f.Name())
			fatalIf(err)
		}()
		defer tmp.Close()
		defer f.Close()

		reader := replace.NewReader(f, flag.Arg(0), flag.Arg(1))
		_, err = io.Copy(tmp, reader)
		if err != nil {
			os.Remove(tmp.Name())
			fatalIf(err)
		}
	default:
		fmt.Fprintf(os.Stderr, "usage %s: [from] [to] {file}\n", os.Args[0])
		os.Exit(1)
	}
}
