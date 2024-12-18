package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	timephrase "github.com/tmichov/TimePhrase"
	"github.com/tmichov/TimeneyeCLI/cmd/request"
)

type EntryRequest struct {
	ProjectID int    `json:"project_id"`
	PhaseID   int    `json:"phase_id"`
	Minutes   int    `json:"minutes"`
	Date      string `json:"date"`
	Notes     string `json:"notes"`
}

func Create(command string, args []string) error {
	arguments, err := parseArgs(command, args)
	if err != nil {
		return err
	}

	if _, ok := arguments["help"]; ok {
		var builder strings.Builder

		builder.WriteString("Description\n")
		builder.WriteString("  - Create a new entry/record\n")
		builder.WriteString("\n")
		builder.WriteString("Usage\n")
		builder.WriteString("  - timeneye create entry -flag <data>\n")
		builder.WriteString("\n")
		builder.WriteString("Options\n")
		builder.WriteString("  -t, --type <type>               Type of the entry. Can be 'project', 'phase', or 'entry'. Default is 'entry'\n")
		builder.WriteString("  -p, --project <project>         Project of the entry. Required if creating an entry or phase.\n")
		builder.WriteString("  -d, --date <date>               Date of the entry. Required if creating an entry.\n")
		builder.WriteString("  -n, --name <name>               Name of the entry. Required.\n")
		builder.WriteString("  -l, --duration <duration>       Duration of the entry in hours. Required if creating an entry.\n")
		builder.WriteString("  -D, --description <description> Description of the entry. Required if creating an entry.\n")
		builder.WriteString("\n")
		builder.WriteString("Examples\n")
		builder.WriteString("  Create an entry:\n")
		builder.WriteString("  - timeneye create -t entry -d yesterday -p project -n name -l 1 -D description\n")
		builder.WriteString("  Create a project:\n")
		builder.WriteString("  - timeneye create -t project -n name \n")
		builder.WriteString("  Create a project phase:\n")
		builder.WriteString("  - timeneye create -t phase -n name -p project \n")

		fmt.Println(builder.String())
		return nil
	}

	cType, ok := arguments["type"]
	if !ok {
		cType = "entry"
	}

	switch cType {
	case "entry":
		entry(command, arguments)
	case "project":
		fmt.Println("Creating project")
	case "phase":
		fmt.Println("Creating phase")
	default:
		fmt.Println("Unknown type")
		return nil
	}

	return nil
}

func entry(command string, args map[string]string) {
	d, ok := args["date"]
	if !ok {
		fmt.Println("Date is required")
		return
	}

	date, err := timephrase.Parse(d, time.Now())
	if err != nil {
		fmt.Println("Err parsing date", err, date)
		return
	}

	project, ok := args["project"]
	if !ok {
		fmt.Println("Project is required")
		return
	}

	durationStr, ok := args["duration"]
	if !ok {
		fmt.Println("Duration is required")
		return
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		fmt.Println("Err parsing duration", err)
		return
	}

	description, ok := args["description"]
	if !ok {
		fmt.Println("Description is required")
		return
	}

	token, err := getToken()
	if err != nil {
		fmt.Println("Error getting token", err)
		return
	}

	projects, err := fetchProjects(token)
	if err != nil {
		fmt.Println("Error fetching projects", err)
		return
	}

	projectId, err := getProjectId(projects, project)
	phaseId, err := getPhaseId(projectId, projects)

	entry := EntryRequest{
		ProjectID: projectId,
		PhaseID:   phaseId,
		Minutes:   duration * 60,
		Date:      date.Format("2006-01-02"),
		Notes:     description,
	}

	err = createEntry(token, entry)
	if err != nil {
		fmt.Println("Error creating entry", err)
		return
	}
}

func getProjectId(projects []Project, project string) (int, error) {
	for _, p := range projects {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(project)) {
			return p.ID, nil
		}
	}

	return 0, fmt.Errorf("Project not found")
}

func getPhaseId(projectId int, projects []Project) (int, error) {
	for _, p := range projects {
		if p.ID == projectId {
			return p.Phases[len(p.Phases)-1].ID, nil
		}
	}
	return 0, nil
}

func createEntry(token string, entry EntryRequest) error {
	_, err := request.Send("POST", "/entries", entry, token)
	if err != nil {
		return err
	}

	fmt.Println("Entry created")
	return nil
}
