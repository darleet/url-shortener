package main

import (
	"fmt"
	"os"
	"shortener/internal/commands"
)

func main() {
	rootCmd, err := commands.InitRunCommand()
	if err != nil {
		fmt.Println("init run command: %w", err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute: %w", err)
		os.Exit(1)
	}
}
