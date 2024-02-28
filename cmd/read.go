package gomovies

import (
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

		if projectID == "" || dataSet == "" || table == "" {
			log.Fatal("Please make sure all env variables are set for GCP (GCP_PROJECT, GCP_DATASET, GCP_TABLE)")
		}

		// Read from BigQuery
		FmtSqlTxt := help.UserInput(dataSet, table)
		c := bigquery.CreateBqClient(projectID)
		bigquery.ReadFromBq(FmtSqlTxt, c)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
