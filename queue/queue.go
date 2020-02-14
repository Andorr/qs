package queue

import (
	"github.com/andorr/qs/utils"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
	"strings"
	"time"
)

type payload struct {
	Exercises []int `json:"exercises"`
	SubjectID int `json:"subjectID"`
	RoomID int `json:"roomID"`
	Desk int `json:"desk"`
	Message string `json:"message"`
	Help bool `json:"help"`
}

func Command() cli.Command {
	return cli.Command {
		Name: "queue",
		Aliases: []string{"q"},
		Subcommands: []cli.Command {
			queueAddCommand(),
			groupCommand(),
			listCommand(),
			removeCommand(),
		},
	}
}

func queueAddCommand() cli.Command {
	return cli.Command{
		Name: "add",
		Description: "Adds the user to the queue",
		Usage: "qs queue add <SUBJECT_NAME> <EXERCISES>",
		Action: handleQueue,
		Flags: []cli.Flag{
			cli.IntFlag{Name: "desk, d", Value: 9},
			cli.IntFlag{Name: "room", Value: 6},
			cli.StringFlag{Name: "group"},
			cli.BoolFlag{Name: "groupIds"},
			cli.BoolFlag{Name: "id"},
			cli.IntFlag{Name: "sleep", Value: 500},
		},
	}
}

// HandleQueue ...
func handleQueue(c *cli.Context) error {
	// Extract arguments
	subjectID, exercises, roomID, deskID, groups, sleep := extractArgParams(c)

	attemptCount := 0
	for {
		queueElementID, err := addQueueElement(&payload {
			Exercises: exercises,
			SubjectID: subjectID,
			RoomID: roomID,
			Desk: deskID,
		})
		if err == nil {
			for _, person := range groups {
				err := AddToGroup(person, queueElementID, exercises)
				if err != nil {
					fmt.Printf(utils.AddToGroupError, err.Error())
				}
			}

			fmt.Printf("Added to queue! QueueElementID: %d", queueElementID)
			break
		}
		attemptCount++
		fmt.Printf("%d -- Error: %s\n", attemptCount, err.Error())
		
		time.Sleep(time.Duration(sleep)*time.Millisecond)
	}
	return nil
}

func addQueueElement(p *payload) (int, error) {
	// Converting request payload to bytes
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err.Error())
	}

	// Preparing response payload
	type resBody struct {
		QueueElementID int `json:"queueElementID"`
	}
	b := &resBody{}

	// Add to queue
	resBytes, err := utils.Execute("res/addQueueElement", bytes)
	if err != nil {
		return -1, err
	}

	// Extract response
	err = json.Unmarshal(resBytes, b)
	if err != nil {
		panic(err.Error())
	}
	return b.QueueElementID, nil
}

func extractArgParams(c *cli.Context) (int, []int, int, int, []int, int) {
	args := c.Args()
	if c.NArg() < 2 {
		log.Fatalf(utils.MissingArguments, c.Command.Usage)
	}

	config := utils.MustGetConfig()

	// Get subject ID
	subjects := config.Subjects
	var subjectID int
	if c.Bool("id") {
		parsedSubjectID, err := strconv.Atoi(args.First())
		if err != nil {
			log.Fatalf(utils.InvalidArgumentInt, args.First())
		}
		subjectID = parsedSubjectID
	} else {
		readSubjectID, ok := subjects[strings.ToLower(args.First())]
		if !ok {
			log.Fatalf(utils.MissingEntity, args.First(), "subjects")
		}
		subjectID = readSubjectID
	}

	// Initialize exercises
	exercises, err := utils.ParseExercises(args.Get(1))
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Parse roomID and deskID
	roomID := c.Int("room")
	deskID := c.Int("desk")
	group := make([]int, 0)

	if c.String("group") != "" {
		people, err := utils.ParsePeople(c.String("group"), c.Bool("groupIds"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		group = people
	}

	fmt.Printf("\nSubjectID: %d\n", subjectID)
	fmt.Printf("Exercises: %+v\n", exercises)
	fmt.Printf("Room id: %d\n", roomID)
	fmt.Printf("Desk id: %d\n", deskID)
	fmt.Printf("Group: %+v\n\n", group)

	return subjectID, exercises, roomID, deskID, group, c.Int("sleep")
}