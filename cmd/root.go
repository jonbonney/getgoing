package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "getgoing",
    Short: "GetGoing is a CLI tool for setting up Go projects using templates",
    Long: `GetGoing is a CLI tool that helps you quickly set up Go projects by using a variety of templates.
You can select from different project types, customize settings, and get started with minimal boilerplate.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to GetGoing! Use `getgoing init` to start a new project.")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
