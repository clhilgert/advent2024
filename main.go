package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/clhilgert/advent2024/pkg/utils"
)

func main() {
	day := flag.Int("d", -1, "Day number (1-25)")
	flag.Parse()
	if *day == -1 {
		fmt.Println("Error: Day flag '-d' must be provided.")
		os.Exit(1)
	}
	utils.FetchInput(*day)
}
