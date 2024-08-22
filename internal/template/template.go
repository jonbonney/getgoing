package template

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
    "sort"

	"github.com/manifoldco/promptui"
    "gopkg.in/yaml.v3"
)

type Variable struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Template struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Variables   []Variable `yaml:"variables"`
	Files       []string   `yaml:"files"`
	DirPath     string     // To store the directory where the template.yaml was found
}


// LoadTemplates scans the given directory for template.yaml files, parses them,
// and returns an array of Template objects.
func LoadTemplates(dir string) ([]Template, error) {
	var templates []Template

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a template.yaml
		if info.IsDir() || info.Name() != "template.yaml" {
			return nil
		}

		// Parse the template.yaml file
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("could not open file %s: %w", path, err)
		}
		defer file.Close()

		var tmpl Template
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&tmpl); err != nil {
			return fmt.Errorf("could not parse file %s: %w", path, err)
		}

		// Store the directory where the template.yaml was found
		tmpl.DirPath = filepath.Dir(path)

		templates = append(templates, tmpl)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not load templates: %w", err)
	}

	// Sort templates by name
	sort.Slice(templates, func(i, j int) bool {
		return templates[i].Name < templates[j].Name
	})

	return templates, nil
}

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
func SelectTemplate(templates []Template) (Template, error) {
    prompt := promptui.Select{
        Label: "Select a Template",
        Items: templates,
        Templates: &promptui.SelectTemplates{
            Label:    "{{ . }}",
            Active:   "> {{ .Name | cyan }}",
            Inactive: "  {{ .Name }}",
            Details: `
----------- Template Details -----------
{{ "Name:" | faint }}        {{ .Name }}
{{ "Description:" | faint }} {{ .Description }}
{{ "Variables:" | faint }}   
{{- range .Variables }}
  - {{ .Name }}: {{ .Description }}
{{- end }}`,
        },
        Size: 10,
    }


    i, _, err := prompt.Run()
    if err != nil {
        return Template{}, fmt.Errorf("could not select template: %w", err)
    }

    return templates[i], nil
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
func GenerateProject(templateDir, targetDir, projectName, moduleName string) error {
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
