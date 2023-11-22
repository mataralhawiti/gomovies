package gomovies

import (
	"github.com/mataralhawiti/gomovies/internal/info"
	"github.com/mataralhawiti/gomovies/internal/parser"
	"github.com/spf13/cobra"
)

const filePath = "resource/sync_movies_full.json"

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "show movies stats",
	Run: func(cmd *cobra.Command, args []string) {
		myMovies := parser.ParseJSON(filePath)
		info.DisplayInfo(myMovies)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
