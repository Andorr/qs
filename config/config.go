package config

import (
	"github.com/andorr/qs/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "config",
		Action: PrintConfig,
		Subcommands: []*cli.Command {
			PeopleCommand(),
			SubjectCommand(),
		},
	}
}

func PrintConfig(c *cli.Context) error {
	// Get config
	config, err := utils.GetOrCreateConfig()
	if err != nil {
		log.Fatalf(utils.ConfigError, err.Error())
	}

	// Print cookie
	fmt.Printf("COOKIE:\n%s\n\n", config.Cookie)

	// Print subjects
	fmt.Printf("SUBJECTS:\n%s\n\n", utils.MapToString(config.Subjects))

	// Print people
	fmt.Printf("PEOPLE:\n%s\n\n", utils.MapToString(config.People))

	return nil
}