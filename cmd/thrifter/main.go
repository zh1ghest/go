package main

import (
	"flag"
	"github.com/zh1ghest/plz.codegen"
	"os"
)

func main() {
	pkgPath := flag.String("pkg", "", "the package to generate generic code for")
	flag.Parse()
	if *pkgPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	wombat.Codegen(*pkgPath)
}
