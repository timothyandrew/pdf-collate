package pdf

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/model"
)

func CollatePDFs(first, second, out string) error {
	writer := model.NewPdfWriter()

	// Set up readers
	readerOne, f, err := model.NewPdfReaderFromFile(first, nil)
	if err != nil {
		return err
	}
	defer f.Close()

	pagesLeft, err := readerOne.GetNumPages()
	if err != nil {
		return err
	}
	fmt.Printf("Loaded file %s with %d pages.\n", first, pagesLeft)

	readerTwo, f, err := model.NewPdfReaderFromFile(second, nil)
	if err != nil {
		return err
	}
	defer f.Close()

	pagesRight, err := readerTwo.GetNumPages()
	if err != nil {
		return err
	}
	fmt.Printf("Loaded file %s with %d pages.\n", second, pagesRight)

	// Determine max page count
	maxPages := 0
	if pagesLeft > maxPages {
		maxPages = pagesLeft
	}
	if pagesRight > maxPages {
		maxPages = pagesRight
	}

	// Interleave
	for i := 1; i <= maxPages; i++ {
		if i <= pagesLeft {
			page, err := readerOne.GetPage(i)
			if err != nil {
				return err
			}

			writer.AddPage(page)
		}

		if i <= pagesRight {
			page, err := readerTwo.GetPage(i)
			if err != nil {
				return err
			}

			writer.AddPage(page)
		}
	}

	f, err = os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()

	err = writer.Write(f)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote file %s\n", out)

	return nil
}
