package cmd

import "fmt"

func Version(command string, args []string) error {
	version := "1.0.0"
	fmt.Printf("Version: %s\n", version)

	return nil
}
