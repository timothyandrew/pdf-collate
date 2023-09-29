/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timothyandrew/pdf-collate/pdf"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of pages in multiple PDFs",
	Run: func(cmd *cobra.Command, args []string) {
		err := pdf.CountPages(args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
