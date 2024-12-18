package cmd

import (
	"fmt"
	"sort"
	"strings"
)

func GetHelp(command string, args []string) error {
	var builder strings.Builder

	builder.WriteString("Usage: timeneye [options]\n")
	builder.WriteString("Options:\n")

	var keys []string
	for k := range Commands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		a := Commands[k]
		builder.WriteString(fmt.Sprintf("  %-10s %s\n", a.Name, a.Description))

		if len(a.Args) == 0 {
			continue
		}
		for _, b := range a.Args {
			builder.WriteString(fmt.Sprintf("      %-6s  %-13s %s\n", b.Flag, b.FlagLong, b.Description))
		}
		if len(a.Args) > 0 {
			builder.WriteString("\n")
		}
	}

	fmt.Println(builder.String())
	return nil
}
