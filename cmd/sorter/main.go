package main

import (
	"LandConnex/cli"
	sorter2 "LandConnex/modules/sorter"
)

func main() {
	strategy := sorter2.NewLastNameSorting()
	parser := sorter2.NewNameParser()
	nameSorter := sorter2.NewNameSorter(strategy, parser)
	cli.Run(nameSorter)
}
