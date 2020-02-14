package qs

import (
	"QS/utils"
	"encoding/json"
)

type Person struct {
	PersonID              int    `json:"personID"`
	SubjectPersonID       int    `json:"subjectPersonID"`
	QueueElementExercises string `json:"queueElementExercises"`
	PersonFirstName       string `json:"personFirstName"`
	PersonLastName        string `json:"personLastName"`
	ExercisesApproved     string `json:"exercisesApproved"`
}

func GetStudentsInQueueElement(queueElementID int) ([]*Person, error){
	bytes, err := queueElementIDPayload(queueElementID)
	if err != nil {
		return nil, err
	}

	res, err := utils.Execute("/res/getStudentsInQueueElement", bytes)
	if err != nil {
		return nil, err
	}

	students := make([]*Person, 0)
	err = json.Unmarshal(res, &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func queueElementIDPayload(queueElementID int) ([]byte, error) {
	// Define and initialize payload
	var payload struct {
		QueueElementID int `json:"queueElementID"`
	}
	payload.QueueElementID = queueElementID

	// Convert payload to json
	return json.Marshal(&payload)
}