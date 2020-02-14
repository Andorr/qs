package queue

import (
	"github.com/andorr/qs/qs"
	"github.com/andorr/qs/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"strconv"
)

func removeCommand() *cli.Command {
	return &cli.Command{
		Name: "remove",
		Usage: "qs queue remove <SUBJECT_NAME> <QUEUE_ELEMENT_ID>",
		Action: handleRemove,
		Flags: []cli.Flag {
			&cli.BoolFlag{Name: "id"},
		},
	}
}

func handleRemove(ctx *cli.Context) error {
	config := utils.MustGetConfig()

	// Get queue element id
	if ctx.NArg() < 1 {
		log.Fatalf(utils.MissingArguments, ctx.Command.Usage)
	}

	// Get subject ID
	subjectIDRaw := ctx.Args().First()
	subjectID := parseSubjectID(ctx, subjectIDRaw, config)

	// Get queue element ID
	queueElementIDRaw := ctx.Args().Get(1)
	queueElementID, err := strconv.Atoi(queueElementIDRaw)
	if err != nil {
		log.Fatalf(utils.InvalidArgumentInt, queueElementIDRaw)
	}

	err = qs.RemoveFromQueue(&qs.DeleteQueueElement{
		SubjectID:      subjectID,
		QueueElementID: queueElementID,
	})

	if err != nil {
		log.Fatalf(utils.QueueRemoveError, queueElementID, err.Error())
	}

	fmt.Println("Done!")

	return nil
}