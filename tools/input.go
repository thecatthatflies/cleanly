package tools

import (
    "flag"
)

type Input struct {
    Path       string
    Categories []string
    NoClean    bool
    Undo       bool
}

func ParseInput() Input {
    noClean := flag.Bool("no-clean", false, "skip verification")
    // add undo flag here

    flag.Parse()

    path := flag.Arg(0)

    return Input{
        Path:    path,
        NoClean: *noClean,
        Undo:    // your undo flag here,
    }
}