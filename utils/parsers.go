package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseExercises(exercisesString string) ([]int, error) {
	exercises := make([]int, 0)
	exercisesRaw := strings.Split(exercisesString, ",")
	for _, ex := range exercisesRaw {
		exercise, err := strconv.Atoi(ex)
		if err != nil {
			return exercises, fmt.Errorf("Invalid exercise: %s", ex)
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

// ParsePeople parses
func ParsePeople(peopleString string, isIds bool) ([]int, error) {
	config := MustGetConfig()
	people := make([]int, 0)
	peopleToAdd := strings.Split(peopleString, ",")
	knownPeople := config.People
	for _, p := range peopleToAdd {
		var person int
		if isIds {
			id, err := strconv.Atoi(p)
			if err != nil {
				return people, fmt.Errorf("Invalid person Id: %s", p)
			}
			person = id
		} else {
			personId, ok := knownPeople[strings.ToLower(p)]
			if !ok {
				return people, fmt.Errorf("Person with id %s is not registered!\n", p)
			}
			person = personId
		}
		people = append(people, person)
	}
	return people, nil
}

// Distinct ...
func Distinct(slice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}