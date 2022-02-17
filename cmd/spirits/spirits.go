package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	command := os.Args[1]
	var f func() error
	switch command {
	case "battle":
		f = runBattle
	case "server":
		f = runServer
	default:
		fmt.Printf("error: unknown command: %s\n", command)
		usage()
		os.Exit(1)
	}

	if err := f(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("usage: %s <command> <flags>\n", os.Args[0])
	fmt.Printf("  commands:\n")
	fmt.Printf("    battle\t-\tperform a local battle\n")
	fmt.Printf("    server\t-\trun spirits server\n")
}
