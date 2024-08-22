package template

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

// Helper function to create a mock template.yaml file
func createMockTemplate(t *testing.T, dir string, tmpl Template) {
	t.Helper()

	data, err := yaml.Marshal(&tmpl)
	if err != nil {
		t.Fatalf("Failed to marshal template: %v", err)
	}

	templatePath := filepath.Join(dir, "template.yaml")
	if err := ioutil.WriteFile(templatePath, data, 0644); err != nil {
		t.Fatalf("Failed to write template.yaml: %v", err)
	}
}

func TestLoadTemplates(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := ioutil.TempDir("", "test-load-templates")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Create mock template directories and template.yaml files
	template1 := Template{
		Name:        "Basic Web App",
		Description: "A simple Go web application with a basic HTTP server.",
		Variables: []Variable{
			{Name: "ProjectName", Description: "The name of your project."},
			{Name: "ModuleName", Description: "The Go module name."},
		},
		Files: []string{"main.go", "go.mod"},
	}
	templateDir1 := filepath.Join(tempDir, "webapp")
	if err := os.Mkdir(templateDir1, os.ModePerm); err != nil {
		t.Fatalf("Failed to create mock template directory: %v", err)
	}
	createMockTemplate(t, templateDir1, template1)

	template2 := Template{
		Name:        "CLI Tool",
		Description: "A simple Go CLI tool.",
		Variables: []Variable{
			{Name: "ProjectName", Description: "The name of your project."},
			{Name: "ModuleName", Description: "The Go module name."},
		},
		Files: []string{"main.go", "go.mod"},
	}
	templateDir2 := filepath.Join(tempDir, "clitool")
	if err := os.Mkdir(templateDir2, os.ModePerm); err != nil {
		t.Fatalf("Failed to create mock template directory: %v", err)
	}
	createMockTemplate(t, templateDir2, template2)

	// Test the LoadTemplates function
	templates, err := LoadTemplates(tempDir)
	if err != nil {
		t.Fatalf("LoadTemplates failed: %v", err)
	}

	// Verify that the correct number of templates were loaded
	if len(templates) != 2 {
		t.Fatalf("Expected 2 templates, got %d", len(templates))
	}

	// Verify that the loaded templates match the expected data
	expectedTemplates := []Template{template1, template2}
	for i, tmpl := range templates {
		if tmpl.Name != expectedTemplates[i].Name {
			t.Errorf("Expected template name %s, got %s", expectedTemplates[i].Name, tmpl.Name)
		}
		if tmpl.Description != expectedTemplates[i].Description {
			t.Errorf("Expected template description %s, got %s", expectedTemplates[i].Description, tmpl.Description)
		}
		if len(tmpl.Variables) != len(expectedTemplates[i].Variables) {
			t.Errorf("Expected %d variables, got %d", len(expectedTemplates[i].Variables), len(tmpl.Variables))
		}
		if len(tmpl.Files) != len(expectedTemplates[i].Files) {
			t.Errorf("Expected %d files, got %d", len(expectedTemplates[i].Files), len(tmpl.Files))
		}
	}
}
