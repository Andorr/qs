package config

import (
	"github.com/andorr/qs/utils"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
)

func PeopleCommand() cli.Command {
	return cli.Command{
		Name: "people",
		Aliases: []string{"people", "p"},
		Usage: "sql config people",
		Subcommands: []cli.Command {
			cli.Command{
				Name: "list",
				Action: PeopleList,
			},
			cli.Command{
				Name: "add",
				Usage: "qs config people add <NAME> <PERSON_ID>",
				Action: PeopleCreate,
			},
			cli.Command{
				Name: "remove",
				Usage: "qs config people remove <NAME>",
				Action: PeopleRemove,
			},
		},
	}
}

func PeopleCreate(c *cli.Context) error {
	if c.NArg() < 2 {
		log.Fatalf(utils.MissingArguments, c.Command.Usage)
	}

	// Get necessary parameters
	args := c.Args()
	personName := args.First()
	personIDRaw := args.Get(1)
	personID, err := strconv.Atoi(personIDRaw)
	if err != nil {
		log.Fatalf(utils.InvalidArgumentInt, personIDRaw)
	}

	// Get config
	config := utils.MustGetConfig()
	people := config.People

	// Check if the person already exists
	if _, ok := people[personName]; ok {
		log.Fatalf(utils.AlreadyExists, personName, "the people list")
	}

	people[personName] = personID

	// Save people map
	config.People = people
	utils.MustSaveConfig(config)
	fmt.Println("Done.")

	return nil
}


func PeopleList(c *cli.Context) error {
	config, err := utils.GetOrCreateConfig()
	if err != nil {
		log.Fatalf(utils.ConfigError, err.Error())
	}

	fmt.Println("People list:")
	fmt.Println(utils.MapToString(config.People))

	return nil
}

func PeopleRemove(c *cli.Context) error {
	// Check if number of arguments are included
	if c.NArg() < 1 {
		log.Fatalf(utils.MissingArguments, c.Command.Usage)
	}

	personName := c.Args().First()

	// Delete subject
	config := utils.MustGetConfig()
	people := config.People
	if _, ok := people[personName]; ok {
		delete(people, personName)
	} else {
		log.Fatalf(utils.MissingEntity, personName, "the people list")
	}

	// Save config
	config.People = people
	utils.MustSaveConfig(config)
	fmt.Println("Done.")

	return nil
}
