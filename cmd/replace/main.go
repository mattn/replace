package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mattn/replace"
)

func main() {
	var err error
	flag.Parse()
	f := os.Stdin
	switch flag.NArg() {
	case 2:
	case 3:
		f, err = os.Open(flag.Arg(2))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	default:
		fmt.Fprintf(os.Stderr, "usage %s: [from] [to] {file}\n", os.Args[0])
		os.Exit(1)
	}

	_, err = io.Copy(replace.NewWriter(os.Stdout, flag.Arg(0), flag.Arg(1)), f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}
