package main

import (
	"fmt"
	"os"

	"github.com/ankeesler/spirits/internal/clientcli"
)

func main() {
	if err := clientcli.Run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
