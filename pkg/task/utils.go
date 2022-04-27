package task

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/rs/zerolog/log"
)

func displayInfo(info *buildInfo) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "src"},
			{Align: simpletable.AlignCenter, Text: "abi"},
			{Align: simpletable.AlignCenter, Text: "bin"},
		},
	}

	for name, paths := range info.contracts {
		r := []*simpletable.Cell{
			{Text: name},
		}

		for _, path := range paths {
			r = append(
				r,
				&simpletable.Cell{
					Text: fmt.Sprintf(`%s`, path),
				},
			)
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())

	log.Info().Msg("contracts compiled successfully")
}
