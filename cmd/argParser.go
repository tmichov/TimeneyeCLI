package cmd

import "strings"

func parseArgs(command string, args []string) (map[string]string, error) {
	data := make(map[string]string)
	var cmdArgs []Arg

	for _, cmd := range Commands {
		if cmd.Name == command {
			cmdArgs = cmd.Args
			break
		}
	}

	selected := ""
	for _, arg := range args {
		if arg[0] == '-' {
			selected = ""
			for _, cmdArg := range cmdArgs {
				if cmdArg.Flag == arg {
					selected = cmdArg.Name
					data[selected] = ""
					break
				}
			}
		} else {
			if selected == "" {
				continue
			}
			curr, ok := data[selected]
			if ok {
				curr += " "
			}
			data[selected] = strings.Trim(curr+arg, " ")
		}
	}

	return data, nil
}
