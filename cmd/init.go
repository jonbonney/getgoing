package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonbonney/getgoing/internal/template"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize a new Go project using a template",
    Run: func(cmd *cobra.Command, args []string) {
        repoURL, _ := cmd.Flags().GetString("template-repo")

        // Clone the remote template repository
        tempDir, err := template.CloneRepository(repoURL)
        if err != nil {
            fmt.Printf("Error cloning repository: %v\n", err)
            os.Exit(1)
        }
        defer os.RemoveAll(tempDir) // Clean up after we're done

        // Load available templates
        templates, err := template.LoadTemplates(tempDir)
        if err != nil {
            fmt.Printf("Error loading templates: %v\n", err)
            os.Exit(1)
        }

        // Prompt the user to select a template
        selectedTemplate, err := template.SelectTemplate(templates)
        if err != nil {
            fmt.Printf("Error selecting template: %v\n", err)
            os.Exit(1)
        }

        // Prompt the user for project details (e.g., name, module)
        projectName, moduleName, err := template.GetProjectDetails()
        if err != nil {
            fmt.Printf("Error getting project details: %v\n", err)
            os.Exit(1)
        }

        // Generate the project
        targetDir := filepath.Join(".", projectName)
        err = template.GenerateProject(selectedTemplate.DirPath, targetDir, projectName, moduleName)
        if err != nil {
            fmt.Printf("Error generating project: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("Project %s initialized successfully!\n", projectName)
    },
}


func init() {
    rootCmd.AddCommand(initCmd)
    initCmd.Flags().StringP("template-repo", "r", "https://github.com/jonbonney/getgoing-templates", "URL of the template repository")
}
