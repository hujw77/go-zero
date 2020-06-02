package main

import (
	"flag"
	"fmt"
	"os"

	"git.atmatrix.org/k12/zero/tool/protobuf/pkg/gen"
	"git.atmatrix.org/k12/zero/tool/protobuf/pkg/generator"
	bmgen "git.atmatrix.org/k12/zero/tool/protobuf/protoc-gen-bm/generator"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(generator.Version)
		os.Exit(0)
	}

	g := bmgen.BmGenerator()
	gen.Main(g)
}
