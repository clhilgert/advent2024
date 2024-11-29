package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/clhilgert/advent2024/pkg/utils"
)

func main() {
	day := flag.Int("n", -1, "Day number (1-25)")
	year := flag.Int("y", 2024, "Year")
	outFile := flag.String("o", "input.txt", "Filename")
	outDir := flag.String("dir", "day%02d", "Output directory")
	flag.Parse()

	if *day == -1 {
		fmt.Println("Error: Day flag '-n' must be provided.")
		os.Exit(1)
	}

	var formattedOutDir string
	if *outDir == "day%02d" {
		formattedOutDir = fmt.Sprintf(*outDir, *day)
	} else {
		formattedOutDir = *outDir
	}

	utils.FetchInput(*day, *year, *outFile, formattedOutDir)
}
