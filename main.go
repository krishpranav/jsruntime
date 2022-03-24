package main

import (
	"fmt"
	"os"

	"github.com/krishpranav/jsruntime/cmd"
)

func main() {
	program := cmd.Execute()

	if err := program.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
