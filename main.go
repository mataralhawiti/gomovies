package main

import (
	"flag"

	"github.com/mataralhawiti/gomovies/internal" // the directory Not the file!
	"github.com/mataralhawiti/gomovies/parser"   // the directory Not the file!
)

var (
	// Flags
	runMod int
	filePath string
	displayStats bool
	PythonBinPath string
)

func main() {
	flag.IntVar(&runMod, "mod", 0, "0 to display stats about your provided movies json or 1 to run python scrtip to generate movies json file. Default: 0")
	flag.StringVar(&filePath, "fp", ".", "movies JSON file path")
	flag.BoolVar(&displayStats, "st", true, "Display stats about my movies. Default: true")
	flag.StringVar(&PythonBinPath, "py", "python", "Your local machine python bin path. Default python")
	// parse all flags
	flag.Parse()

	if runMod == 1 {
		// exect python script
		parser.RunPyScript(PythonBinPath)
		return
	}

	// parse movies json file
	myMovies := parser.ParseJSON(filePath)

	// New imple - display stats about your movies list
	if displayStats {
		internal.DisplayInfo(myMovies)
	}
}
