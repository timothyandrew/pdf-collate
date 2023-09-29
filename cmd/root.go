/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/timothyandrew/pdf-collate/pdf"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pdf-collate [flags] FIRST_FILE SECOND_FILE",
	Short: "Merge two PDFs (duplex scanning with a single-sided scanner)",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := cmd.Flags().GetString("output")
		err := pdf.CollatePDFs(args[0], args[1], out)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("output", "o", "", "Path/filename for the output PDF")
}
