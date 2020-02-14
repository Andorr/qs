package config

import (
	"github.com/andorr/qs/utils"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
)

func SubjectCommand() cli.Command {
	return cli.Command{
		Name: "subject",
		Aliases: []string{"subjects", "s"},
		Usage: "sql config subject",
		Subcommands: []cli.Command {
			cli.Command{
				Name: "list",
				Action: SubjectList,
			},
			cli.Command{
				Name: "add",
				Usage: "qs config subjects add <NAME> <SUBJECT_ID>",
				Action: SubjectCreate,
			},
			cli.Command{
				Name: "remove",
				Usage: "qs config subjects remove <NAME>",
				Action: SubjectRemove,
			},
		},
	}
}

func SubjectCreate(c *cli.Context) error {
	if c.NArg() < 2 {
		log.Fatalf(utils.MissingArguments, c.Command.Usage)
	}

	// Get necessary parameters
	args := c.Args()
	subjectName := args.First()
	subjectIDRaw := args.Get(1)
	subjectID, err := strconv.Atoi(subjectIDRaw)
	if err != nil {
		log.Fatalf(utils.InvalidArgumentInt, subjectIDRaw)
	}

	// Get config
	config := utils.MustGetConfig()
	subjects := config.Subjects

	// Check if it subject already exists
	if _, ok := subjects[subjectName]; ok {
		log.Fatalf(utils.AlreadyExists, subjectName, "subjects")
	}

	subjects[subjectName] = subjectID

	// Save subject
	config.Subjects = subjects
	utils.MustSaveConfig(config)
	fmt.Println("Done.")

	return nil
}

func SubjectList(c *cli.Context) error {
	config := utils.MustGetConfig()

	fmt.Println("Subject list:")
	fmt.Println(utils.MapToString(config.Subjects))

	return nil
}

func SubjectRemove(c *cli.Context) error {
	// Check if number of arguments are included
	if c.NArg() < 1 {
		log.Fatalf(utils.MissingArguments, c.Command.Usage)
	}

	subjectName := c.Args().First()

	// Delete subject
	config := utils.MustGetConfig()
	subjects := config.Subjects
	if _, ok := subjects[subjectName]; ok {
		delete(subjects, subjectName)
	} else {
		log.Fatalf(utils.MissingEntity, subjectName, "subjects")
	}

	// Save config
	config.Subjects = subjects
	utils.MustSaveConfig(config)
	fmt.Println("Done.")

	return nil
}