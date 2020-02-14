package queue

import (
	"github.com/andorr/qs/utils"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
)

type groupPayload struct {
	Exercises []int `json:"exercises"`
	QueueElementID int `json:"queueElementID"`
	SubjectPersonID int `json:"subjectPersonID"`
}

func groupCommand() cli.Command {
	return cli.Command {
		Name: "group",
		Aliases: []string{"g"},
		Description: "Sets a person in group with an existing queueElement.",
		Usage: "qs queue group <QUEUE_ELEMENT_ID> <NAMES> <EXERCISES>",
		Action: handleGroup,
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "ids"},
		},
	}
}

// HandleGroup ...
func handleGroup(c *cli.Context) {
	// Check if number of arguments is valid
	if c.NArg() < 3 {
		log.Fatalf(utils.MissingArguments, c.Command.Usage)
	}
	args := c.Args()

	// Get queueElementID
	queueElementID, err := strconv.Atoi(args.First())
	if err != nil {
		log.Fatalf("Given queueElementID is invalid! %s", args.First())
	}

	// Parsing groups
	group, err := utils.ParsePeople(args.Get(1), c.Bool("groupIds"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	exercises, err := utils.ParseExercises(args.Get(2))
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, person := range group {
		err := AddToGroup(person, queueElementID, exercises)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
}

// AddToGroup attempts to add a person to a group
func AddToGroup(personID int, queueElementID int, exercises []int) error {
	payload := &groupPayload {
		SubjectPersonID: personID,
		QueueElementID: queueElementID,
		Exercises: exercises,
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = utils.Execute("res/addPersonToQueueElement", bytes)
	return err
}