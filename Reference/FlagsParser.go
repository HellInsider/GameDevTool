package main

import (
	"flag"
)

type Flags struct {
	RPM    int
	Offset int
	Limit  int
}

func ParseFlags() Flags {
	flags := Flags{}
	flag.IntVar(&flags.RPM, "rpm", 30, "Request per minute")
	flag.IntVar(&flags.Offset, "offset", 0, "Update from game serial number X")
	flag.IntVar(&flags.Limit, "limit", 0, "update untill X")
	flag.Parse()

	return flags
}
