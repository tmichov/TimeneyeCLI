package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tmichov/TimeneyeCLI/config"
)

type TokenConfig struct {
	Token string `json:"token"`
}

func AuthToken(command string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Auth token requires -t flag with token value. Run '-h' for help.")
	}

	arguments, err := parseArgs(command, args)
	if err != nil {
		return err
	}

	if len(arguments) == 0 {
		return fmt.Errorf("Auth token requires -t flag with token value. Run '-h' for help.")
	}

	_, ok := arguments["help"]
	if ok {
		var builder strings.Builder

		builder.WriteString("Description:\n")
		builder.WriteString("  - Set the personal access token to authenticate all requests.\n")
		builder.WriteString("  - Go to https://app.timeneye.com/developers for more info on how to set up a personal access token.\n")

		builder.WriteString("\n")
		builder.WriteString("Usage: auth token -t <token>\n")
		builder.WriteString("\n")
		builder.WriteString("Options:\n")
		builder.WriteString("  -t, --token <token>  Token value\n")
		builder.WriteString("  -h, --help           Show help for auth token\n")

		builder.WriteString("\n")
		builder.WriteString("Examples:\n")
		builder.WriteString("  auth token -t 1234567890abcdef\n")

		fmt.Println(builder.String())
		return nil
	}

	token := TokenConfig{
		Token: arguments["token"],
	}

	err = config.WriteConfig("config/token.json", token)
	if err != nil {
		return err
	}

	return nil
}

func getToken() (string, error) {
	tokenStr, err := config.ReadConfig("config/token.json")
	if err != nil {
		return "", err
	}

	var tokenConfig TokenConfig
	if err := json.Unmarshal(tokenStr, &tokenConfig); err != nil {
		return "", fmt.Errorf("Error while unmarshalling token config")
	}

	if tokenConfig.Token == "" {
		var builder strings.Builder

		builder.WriteString("Before using this command you need to authenticate yourself.\n")
		builder.WriteString("Please use the command 'timeneye auth' to authenticate yourself.\n")
		builder.WriteString("To see more details use 'timeneye auth -h'\n")

		fmt.Println(builder.String())
		return "", fmt.Errorf("")
	}

	return tokenConfig.Token, nil
}
