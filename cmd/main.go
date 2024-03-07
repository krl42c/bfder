package main

import (
	"bfder/cmd/bfder"
	"fmt"
	"os"
)

const HELP = "Usage: ./fder [search_term] [directory]\n"
const VERSION = "0.1"

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("fder%s\n%s", VERSION, HELP)
		return
	}

	if len(args) < 3 {
		fmt.Println("Not enough arguments provided")
		fmt.Print(HELP)
		return
	}

	search := args[1]
	dir := args[2]

	if _, err := os.Stat(dir); err != nil {
		fmt.Printf("%s is not a directory\n", dir)
		fmt.Print(HELP)
		return
	}

	bfder.Run(search, dir)
}

func PrintHelp() {
	fmt.Printf("fder\n")
}
