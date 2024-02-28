package gomovies

import (
	"fmt"
	"log"
	"os"

	"github.com/mataralhawiti/gomovies/internal/bigquery" // the directory Not the file!
	"github.com/mataralhawiti/gomovies/internal/help"     // the directory Not the file!
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read pre-defined queries from GCP BigQuery",
	Run: func(cmd *cobra.Command, args []string) {
		// OS env
		projectID := os.Getenv("GCP_PROJECT")
		dataSet := os.Getenv("GCP_DATASET")
		table := os.Getenv("GCP_TABLE")

		// Read from BigQuery
		sqlTxt := help.UserInput()
		FmtSqlTxt := fmt.Sprintf(sqlTxt, dataSet, table)

		if projectID == "" {
			log.Fatal("- GCP_PROJECT environment variable must be set.\n# bash: export GCP_PROJECT=xxx\n# powershell: $env:GCP_PROJECT=xxx")
		}
		c := bigquery.CreateBqClient(projectID)
		bigquery.ReadFromBq(FmtSqlTxt, c)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
