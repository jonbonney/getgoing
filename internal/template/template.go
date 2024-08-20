package template

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

// CloneRepository clones the remote template repository to a temporary directory
func CloneRepository(repoURL string) (string, error) {
    tempDir, err := os.MkdirTemp("", "getgoing-")
    if err != nil {
        return "", fmt.Errorf("could not create temp directory: %w", err)
    }

    cmd := exec.Command("git", "clone", repoURL, tempDir)
    if err := cmd.Run(); err != nil {
        return "", fmt.Errorf("could not clone repository: %w", err)
    }

    return tempDir, nil
}

// ListTemplates lists available templates in the repository
func ListTemplates(repoDir string) ([]string, error) {
    files, err := os.ReadDir(repoDir)
    if err != nil {
        return nil, fmt.Errorf("could not list templates: %w", err)
    }

    var templates []string
    for _, file := range files {
        if file.IsDir() {
            templates = append(templates, file.Name())
        }
    }

    return templates, nil
}

// SelectTemplate prompts the user to select a template
func SelectTemplate(templates []string) (string, error) {
    prompt := promptui.Select{
        Label: "Select a Template",
        Items: templates,
    }

    _, result, err := prompt.Run()
    if err != nil {
        return "", fmt.Errorf("could not select template: %w", err)
    }

    return result, nil
}

// GetProjectDetails prompts the user for project details
func GetProjectDetails() (string, string, error) {
    prompt := promptui.Prompt{
        Label: "Enter Project Name",
    }
    projectName, err := prompt.Run()
    if err != nil {
        return "", "", fmt.Errorf("could not get project name: %w", err)
    }

    prompt = promptui.Prompt{
        Label: "Enter Module Name",
        Default: fmt.Sprintf("github.com/yourusername/%s", projectName),
    }
    moduleName, err := prompt.Run()
    if err != nil {
        return "", "", fmt.Errorf("could not get module name: %w", err)
    }

    return projectName, moduleName, nil
}

// GenerateProject generates the project in the target directory
func GenerateProject(templateName, repoDir, targetDir, projectName, moduleName string) error {
    templateDir := filepath.Join(repoDir, templateName)

    err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        relPath := strings.TrimPrefix(path, templateDir)
        relPath = strings.TrimPrefix(relPath, string(filepath.Separator))

        targetPath := filepath.Join(targetDir, relPath)

        if info.IsDir() {
            return os.MkdirAll(targetPath, os.ModePerm)
        }

        content, err := os.ReadFile(path)
        if err != nil {
            return err
        }

        replacedContent := strings.ReplaceAll(string(content), "{{ProjectName}}", projectName)
        replacedContent = strings.ReplaceAll(replacedContent, "{{ModuleName}}", moduleName)

        return os.WriteFile(targetPath, []byte(replacedContent), info.Mode())
    })

    if err != nil {
        return fmt.Errorf("could not generate project: %w", err)
    }

    return nil
}
