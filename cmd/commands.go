package cmd

type Arg struct {
	Name        string
	Flag        string
	FlagLong    string
	Value       string
	Description string
}

type Command struct {
	Name        string
	Args        []Arg
	Description string
	Executor    func(command string, args []string) error
}

var Commands = make(map[string]Command)

func addCommand(name, description string, executor func(command string, args []string) error) {
	Commands[name] = Command{
		Name:        name,
		Description: description,
		Executor:    executor,
	}
	addCommandArgument(name, "help", "-h", "--help", "string", "Display help for this command")
}

func addCommandArgument(name, argName, flag, flagLong, value, description string) {
	command := Commands[name]
	command.Args = append(command.Args, Arg{
		Name:        argName,
		Flag:        flag,
		FlagLong:    flagLong,
		Value:       value,
		Description: description,
	})
	Commands[name] = command
}

func SetupCommands() error {
	addCommand("auth", "Authenticate to the API", AuthToken)
	addCommandArgument("auth", "token", "-t", "--token", "string", "Token to authenticate")

	addCommand("projects", "List all projects", Projects)

	addCommand("create", "Create an entity", Create)
	addCommandArgument("create", "type", "-t", "--type", "string", "Type of entity to create")
	addCommandArgument("create", "project", "-p", "--project", "string", "Project of the entity to create")
	addCommandArgument("create", "date", "-d", "--date", "string", "Date of the entity to create")
	addCommandArgument("create", "name", "-n", "--name", "string", "Name of the entity to create")
	addCommandArgument("create", "duration", "-l", "--duration", "string", "Duration of the entity to create")
	addCommandArgument("create", "description", "-D", "--description", "string", "Description of the entity to create")

	addCommand("help", "Get help", GetHelp)

	addCommand("version", "Get the version of the CLI", Version)

	return nil
}
