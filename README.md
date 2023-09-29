# pdf-collate

Merge two PDFs page-by-page, useful when collating PDFs of a double-sided document scanned with a single-sided ADF scanner.

## Usage

```bash
$ go build
$ ./pdf-collate -o out.pdf ~/Downloads/{1,2}.pdf
```