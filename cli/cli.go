package cli

import (
	"LandConnex/sorter"
	"fmt"
	cli "github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"strings"
)

func Run(appSorter sorter.Sorter) {
	app := &cli.App{
		Name:  "name-sorter",
		Usage: "Sorts a list of names from a file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "File containing names to sort",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "sort-by",
				Usage: "Sort by 'first' or 'last' name",
				Value: "last",
			},
			&cli.BoolFlag{
				Name:  "desc",
				Usage: "Sort in descending order",
				Value: false,
			},
		},
		Action: func(c *cli.Context) error {
			fileName := c.String("file")
			sortBy := c.String("sort-by")
			desc := c.Bool("desc")

			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return err
			}

			names := strings.Split(strings.TrimSpace(string(content)), "\n")

			if sortBy == "first" {
				appSorter = sorter.NewNameSorter(&sorter.FirstNameSorting{}, sorter.NewNameParser())
			} else {
				appSorter = sorter.NewNameSorter(&sorter.LastNameSorting{}, sorter.NewNameParser())
			}

			sortedNames := appSorter.Sort(names, desc)

			outputFileName := fileName[:len(fileName)-len(".txt")] + "-sorted.txt"
			err = ioutil.WriteFile(outputFileName, []byte(strings.Join(sortedNames, "\n")), 0644)
			if err != nil {
				fmt.Printf("Error writing sorted file: %v\n", err)
				return err
			}

			fmt.Printf("Sorted names written to %s\n", outputFileName)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
