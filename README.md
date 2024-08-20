# GetGoing

**GetGoing** is a highly extensible CLI tool for quickly starting new Go projects using a variety of templates. It streamlines the project setup process by allowing you to select from different types of Go project templates, customize key settings, and get started with minimal boilerplate.

## Features

- **Template-based Project Initialization:** Choose from various templates like web apps, CLI tools, microservices, and more.
- **Customizable Project Settings:** Configure your project name, module name, and other settings through an interactive wizard.
- **Extensibility:** Easily extend the tool by adding new templates or modifying existing ones.

## Installation

You can install GetGoing using go install:
```
go install github.com/jonbonney/getgoing@latest
```
Make sure your GOPATH/bin is added to your PATH so you can run getgoing from anywhere.

## Usage

### Initializing a New Project

To initialize a new Go project using GetGoing, run:
```
getgoing init
```
You will be prompted to select a project type and template, enter your project details, and the tool will generate the project for you.

### Example

1. **Run the init command:**
```
getgoing init
```
2. **Follow the prompts:**
- Select a template (e.g., "Basic Web App").
- Enter your project name (e.g., my-web-app).
- Confirm the Go module name (e.g., github.com/yourusername/my-web-app).

3. **Navigate to your new project:**
```
cd my-web-app
```
4. **Run your project:**
```
go run main.go
```
You should see the output `Starting server on :8080`. Visit `http://localhost:8080` to see your running web app.

## To-Do List

- Expand the template repo config variable into a template repo index to more easily add personal repos and have multiple at once.
- Optionally cache templates locally.
- Implement update-templates argument to update the cache.
- Search recursively through template directory to find the template.yaml files and parse them to display template description and other info to the user in the selection menu, or just include that information in the newly created index.
- Add more templates.
- Add tags to template.yaml so that we can first allow the user to select a tag group before being shown the full list of templates (not high priority until we have more templates).

## Contributing

Contributions are welcome! Here's how you can help:

1. **Fork the Repository:**
- Fork the getgoing-templates repository to add or update templates.

2. **Add or Update Templates:**
- Create a new directory for your template.
- Add a template.yaml file describing your template.
- Include all necessary files (e.g., main.go, go.mod).

3. **Submit a Pull Request:**
- After making your changes, submit a pull request to the main repository.

4. **Documentation:**
- Ensure that your template is well-documented, including descriptions in template.yaml and example usage if applicable.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
