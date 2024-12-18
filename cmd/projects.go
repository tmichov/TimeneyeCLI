package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tmichov/TimeneyeCLI/cmd/request"
	"github.com/tmichov/TimeneyeCLI/config"
)

type Payload struct {
	Load []string `json:"load"`
}

type Phase struct {
	ID              int    `json:"phase_id"`
	ProjectID       int    `json:"project_id"`
	Name            string `json:"name"`
	IsOpen          int    `json:"is_open"`
	PhaseCategoryID int    `json:"phase_category_id"`
	BudgetMinutes   int    `json:"budget_minutes"`
}

type Project struct {
	Name   string  `json:"name"`
	ID     int     `json:"project_id"`
	Phases []Phase `json:"phases"`
}

type Resp struct {
	Data []Project `json:"data"`
}

func Projects(command string, args []string) error {
	arguments, err := parseArgs(command, args)
	if err != nil {
		return err
	}
	_, ok := arguments["help"]
	if ok {
		var builder strings.Builder

		builder.WriteString("Description:\n")
		builder.WriteString("  - List all projects and their phases\n")

		builder.WriteString("\n")
		builder.WriteString("Usage: `timeneye projects`\n")
		builder.WriteString("\n")
		builder.WriteString("Requirements:\n")
		builder.WriteString("  - You need to set the authentication token before using this command.\n")
		builder.WriteString("  - See `timeneye auth -h` for more details.\n")
		builder.WriteString("\n")
		builder.WriteString("Options:\n")
		builder.WriteString("  -h, --help           Show help for this command\n")

		builder.WriteString("\n")
		builder.WriteString("Examples:\n")
		builder.WriteString("  projects\n")

		fmt.Println(builder.String())
		return nil
	}

	token, err := getToken()
	if err != nil {
		return nil
	}

	projects, err := fetchProjects(token)
	if err != nil {
		return err
	}

	printProjects(projects)

	err = config.WriteConfig("config/projects.json", projects)
	if err != nil {
		return err
	}

	return nil
}

func fetchProjects(token string) ([]Project, error) {
	currentStr, err := config.ReadConfig("config/projects.json")
	if err != nil {
		return []Project{}, err
	}

	currentProjects := []Project{}
	if err := json.Unmarshal(currentStr, &currentProjects); err != nil {
		return []Project{}, err
	}

	if len(currentProjects) > 0 {
		return currentProjects, nil
	}

	payload := Payload{Load: []string{"phases"}}

	data, err := request.Send("GET", "projects", payload, token)
	if err != nil {
		return nil, err
	}

	projectList := []Project{}
	if err := json.Unmarshal(data, &projectList); err != nil {
		fmt.Println("Error parsing response", err)
	}

	return projectList, nil
}

func printProjects(projects []Project) {
	var builder strings.Builder

	for _, project := range projects {
		builder.WriteString(fmt.Sprintf("ID %d: Name %s\n", project.ID, project.Name))
		for _, phase := range project.Phases {
			builder.WriteString(fmt.Sprintf("    Phase ID %d: Name %s\n", phase.ID, phase.Name))
		}
	}

	fmt.Println(builder.String())
}
