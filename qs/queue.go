package qs

import (
	"QS/utils"
	"encoding/json"
	"time"
)

type QueueElement struct {
	QueueElementID        int         `json:"queueElementID"`
	SubjectID             int         `json:"subjectID"`
	RoomID                int         `json:"roomID"`
	QueueElementDesk      int         `json:"queueElementDesk"`
	QueueElementMessage   interface{} `json:"queueElementMessage"`
	QueueElementPosition  int         `json:"queueElementPosition"`
	QueueElementHelp      int         `json:"queueElementHelp"`
	QueueElementTeacher   int         `json:"queueElementTeacher"`
	QueueElementStartTime time.Time   `json:"queueElementStartTime"`
	OwnerID               int         `json:"ownerID"`
	SubjectPersonID       int         `json:"subjectPersonID"`
	QueueElementExercises string      `json:"queueElementExercises"`
	RoomNumber            string      `json:"roomNumber"`
	RoomImgLink           string      `json:"roomImgLink"`
	PersonFirstName       string      `json:"personFirstName"`
	PersonLastName        string      `json:"personLastName"`
	Groupmembers          string      `json:"groupmembers"`
}

func GetQueue(subjectID int) ([]*QueueElement, error) {
	bytes, err := subjectIDPayload(subjectID)
	if err != nil {
		return nil, err
	}

	res, err := utils.Execute("/res/getQueue", bytes)
	if err != nil {
		return nil, err
	}

	queue := make([]*QueueElement, 0)
	err = json.Unmarshal(res, &queue)
	if err != nil {
		return nil, err
	}
	return queue, nil
}

type Teacher struct {
	SubjectPersonID int    `json:"subjectPersonID"`
	SubjectID       int    `json:"subjectID"`
	SubjectRoleID   int    `json:"subjectRoleID"`
	PersonID        int    `json:"personID"`
	PersonFirstName string `json:"personFirstName"`
	PersonLastName  string `json:"personLastName"`
	PersonEmail     string `json:"personEmail"`
	SystemRoleID    int    `json:"systemRoleID"`
	RoleDescription string `json:"roleDescription"`
}

func GetTeachers(subjectID int) ([]*Teacher, error) {
	bytes, err := subjectIDPayload(subjectID)
	if err != nil {
		return nil, err
	}

	res, err := utils.Execute("/res/regSubjectGetTeachers", bytes)
	teachers := make([]*Teacher, 0)
	err = json.Unmarshal(res, &teachers)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func subjectIDPayload(subjectID int) ([]byte, error) {
	// Define and initialize payload
	var payload struct {
		SubjectID int `json:"subjectID"`
	}
	payload.SubjectID = subjectID

	// Convert payload to json
	return json.Marshal(&payload)
}

type DeleteQueueElement struct {
	SubjectID int `json:"subjectID"`
	QueueElementID int `json:"queueElementID"`
}

func RemoveFromQueue(payload *DeleteQueueElement) error {
	// Parse payload to json
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = utils.Execute("/res/deleteQueueElement", bytes)
	if err != nil {
		return err
	}

	return nil
}