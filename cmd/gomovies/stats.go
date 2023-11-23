package gomovies

import (
	"github.com/mataralhawiti/gomovies/internal/info"
	"github.com/mataralhawiti/gomovies/internal/parser"
	"github.com/spf13/cobra"
)

// var filePath = "resource/sync_movies_full.json"
var filePath string
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "show movies stats",
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		myMovies := parser.ParseJSON(filePath)
		info.DisplayInfo(myMovies)
	},
}

func init() {
	statsCmd.Flags().StringVarP(&filePath, "file_path", "f", "", "Path to movies JSON file. ex: movies.json")
	rootCmd.AddCommand(statsCmd)

	statsCmd.MarkFlagRequired("file_path")
}
