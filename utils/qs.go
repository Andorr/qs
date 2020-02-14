package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://qs.stud.iie.ntnu.no/"

func Execute(endpoint string, payload []byte) ([]byte, error) {
	config := MustGetConfig()
	if config.Cookie == "" {
		log.Fatalf(Unauthenticated)
	}

	req, err:= http.NewRequest("POST", fmt.Sprintf("%s/%s", baseURL, endpoint), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", config.Cookie)
	req.Header.Add("Accept","application/json, text/plain, */*")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Accept-Language","nb-NO,nb;q=0.9,no;q=0.8,nn;q=0.7,en-US;q=0.6,en;q=0.5")
	req.Header.Add("Connection","keep-alive")
	req.Header.Add("Host","qs.stud.iie.ntnu.no")
	req.Header.Add("Origin","https://qs.stud.iie.ntnu.no")
	req.Header.Add("Referer","https://qs.stud.iie.ntnu.no/sTeacher")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil || res == nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Response came with %s", res.Status)
	}

	return ioutil.ReadAll(res.Body)
}