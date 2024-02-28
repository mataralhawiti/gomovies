package gomovies

import (
	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "gomovies",
	Version: version,
	Short:   "CLI for IMDB movies",
	Long:    "CLI to parse IMD movies JSON file, Insert movies info into BigQuery, query movies info from BigQuery",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
