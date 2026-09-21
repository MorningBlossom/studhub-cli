package cmd

import (
	"fmt"
	"os"

	generator "github.com/MorningBlossom/studhub-cli/pkg/studhub-generator"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "service",
	Short: "Generate template files for a new StudHub microservice",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		fmt.Printf("Scaffolding new StudHub service: %s...\n", serviceName)

		if err := generator.GenerateService(serviceName); err != nil {
			fmt.Printf("Failed to generate service: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully generated %s!\n", serviceName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("name", "n", "", "Name of the new service")
}