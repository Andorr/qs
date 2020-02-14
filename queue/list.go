package queue

import (
	"QS/qs"
	"QS/utils"
	"fmt"
	"github.com/inancgumus/screen"
	"github.com/urfave/cli"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func listCommand() cli.Command {
	return cli.Command{
		Name: "list",
		Usage: "qs queue list <SUBJECT_NAME>",
		Action: listQueue,
		Flags: []cli.Flag {
			cli.BoolFlag{Name: "id"},
			cli.IntFlag{Name: "limit", Value: -1},
			cli.IntFlag{Name: "ext"},
		},
	}
}

func listQueue(ctx *cli.Context) {
	config := utils.MustGetConfig()

	// Get subject ID
	if ctx.NArg() < 1 {
		log.Fatalf(utils.MissingArguments, ctx.Command.Usage)
	}
	args := ctx.Args()
	subjectIDRaw := args.First()
	subjectID := parseSubjectID(ctx, subjectIDRaw, config)

	// Get teachers connected to the subject
	teacherMap := map[int]string{}
	teachers, err := qs.GetTeachers(subjectID)
	if err == nil {
		for _, teacher := range teachers {
			teacherMap[teacher.PersonID] = teacher.PersonFirstName + " " + teacher.PersonLastName
		}
	}

	// List queue
	for {
		screen.Clear()
		queueElements, err := qs.GetQueue(subjectID)
		sort.SliceStable(queueElements, func(i, j int) bool {
			return queueElements[i].QueueElementPosition < queueElements[j].QueueElementPosition
		})

		if err != nil {
			log.Fatalf(utils.QueueListError, subjectID, err.Error())
		}

		screen.MoveTopLeft()
		limit := len(queueElements)
		limitFlag := ctx.Int("limit")
		if limitFlag  != -1 {
			limit = limitFlag
		}
		queueElements = queueElements[:limit]
		for i, elem := range queueElements {
			fmt.Println(stringifyQueueElement(i, elem, teacherMap))
		}
		time.Sleep(time.Second*2)
	}
}

func stringifyQueueElement(index int, elem *qs.QueueElement, teachers map[int]string) string {
	sprint := "%-3d: %-40s %-10s %-10s %-10d %-16s %-20s %s"

	// Handle displayName
	displayName := elem.PersonFirstName
	groupMembers := strings.Split(elem.Groupmembers, ",")
	if len(groupMembers) > 1 {
		displayName += fmt.Sprintf(" & Co (%d)", len(groupMembers))
	} else {
		displayName += " " + elem.PersonLastName
	}

	peopleInGroup := ""
	if len(groupMembers) > 1 {
		people, err := qs.GetStudentsInQueueElement(elem.QueueElementID)
		if err == nil {
			for _, person := range people {
				peopleInGroup += person.PersonFirstName + " " + person.PersonLastName + ", "
			}
		}
	}

	// Handle exercises
	exercises := strings.Split(elem.QueueElementExercises, ",")
	exercises = utils.Distinct(exercises)
	exerciseString := strings.Join(exercises, ",")

	queueType := "Godkjenning"
	if elem.QueueElementHelp == 1 {
		queueType = "Hjelp"
	}
	timeSince := time.Since(elem.QueueElementStartTime)
	formattedTime := fmt.Sprintf("%dm%ds", int(timeSince.Minutes()), int(timeSince.Seconds())%60)

	selected := ""
	if elem.QueueElementTeacher != 0 {
		selected = "*****"
		if val, ok := teachers[elem.QueueElementTeacher]; ok {
			selected += " " + val
		}
	}

	output := fmt.Sprintf(sprint,index, displayName, exerciseString, elem.RoomNumber, elem.QueueElementDesk,queueType,formattedTime, selected)
	if len(peopleInGroup) > 0 {
		output += "\n\t\t" + peopleInGroup
	}
	return output
}

func parseSubjectID(ctx *cli.Context, subjectIDRaw string, config *utils.Config) int {
	var subjectID int
	if ctx.Bool("id") {
		id, err := strconv.Atoi(subjectIDRaw)
		if err != nil {
			log.Fatalf(utils.InvalidArgumentInt, subjectIDRaw)
		}
		subjectID = id
	} else {
		if val, ok := config.Subjects[subjectIDRaw]; ok {
			subjectID = val
		} else {
			log.Fatalf(utils.MissingEntity, subjectIDRaw, "subjects")
		}
	}
	return subjectID
}