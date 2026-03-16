package tools

import (
	"flag"
	"strings"
)

type Input struct {
	Path       string
	Categories []string
	NoClean    bool
	Undo       bool
}

func ParseInput() Input {
	noClean := flag.Bool("no-clean", false, "skip cleanup verification after sorting")
	undo := flag.Bool("undo", false, "reverse the last sort")
	c := flag.String("c", "", "only sort specific categories e.g. -c \"images audio\"")

	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	categories := strings.Fields(strings.ToLower(*c))

	return Input{
		Path:       path,
		Categories: categories,
		NoClean:    *noClean,
		Undo:       *undo,
	}
}
