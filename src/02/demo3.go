package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
}

