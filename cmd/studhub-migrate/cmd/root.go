package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "studhub-generate",
	Short: "StudHub CLI for microservices scaffolding",
	Long:  `A command line tool to manage and scaffold infrastructure for the StudHub inter-college platform.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}