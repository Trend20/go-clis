package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apitester",
	Short: "A CLI tool that tests APIs and generates API documentation.",
	Long: `A CLI tool that tests APIs and generates API documentation based on the test results.
Support different formats like OpenAPI/Swagger and include automated testing features.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add additional commands here
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(generateDocCmd)
}
