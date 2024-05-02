package main

import (
	"LandConnex/cli"
	"LandConnex/sorter"
)

func main() {
	strategy := sorter.NewLastNameSorting()
	parser := sorter.NewNameParser()
	nameSorter := sorter.NewNameSorter(strategy, parser)
	cli.Run(nameSorter)
}
