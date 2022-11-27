package main

import (
	"flag"
)

type Flags struct {
	AllGames   bool
	GameDetail bool
	LoadData   bool
}

func ParseFlags() Flags {
	flags := Flags{}
	flag.BoolVar(&flags.AllGames, "allgames", false, "get base info of all games")
	flag.BoolVar(&flags.GameDetail, "gamedetails", false, "get game details")
	flag.BoolVar(&flags.LoadData, "downloaddata", false, "download txt files with games information")
	flag.Parse()

	return flags
}
