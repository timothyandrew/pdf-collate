package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
	"github.com/unidoc/unipdf/v3/model"
)

func CountPages(paths []string) error {
	count := 0

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Filename", "Pages"})

	for _, path := range paths {
		reader, f, err := model.NewPdfReaderFromFile(path, nil)
		if err != nil {
			return err
		}
		defer f.Close()

		pages, err := reader.GetNumPages()
		if err != nil {
			return err
		}

		table.Append([]string{filepath.Base(path), fmt.Sprint(pages)})
		count += pages
	}

	table.SetFooter([]string{"TOTAL", fmt.Sprint(count)})
	table.Render()

	return nil
}
