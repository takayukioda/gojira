package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type CreateSprintParam struct {
	Name          string `json:"name"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	OriginBoardId int64  `json:"originBoardId"`
	Goal          string `json:"goal"`
}

type Sprint struct {
	Id            int64  `json:"id"`
	Self          string `json:"self"`
	State         string `json:"state"`
	Name          string `json:"name"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	CompleteDate  string `json:"completeDate"`
	OriginBoardId int64  `json:"originBoardId"`
	Goal          string `json:"goal"`
}

func (j *Jira) CreateSprint(param CreateSprintParam) (*Sprint, error) {
	uri := fmt.Sprintf("https://%s/rest/agile/1.0/sprint", j.domain)
	buf := bytes.NewBufferString("")
	err := json.NewEncoder(buf).Encode(param)
	if err != nil {
		return nil, err
	}

	sprint := &Sprint{}
	res, err := j.request("POST", uri, buf, sprint)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("Status was not Created: %s\n", res.Status))
	}
	return sprint, nil
}

func (j *Jira) GetSprint(id int64) (*Sprint, error) {
	uri := fmt.Sprintf("https://%s/rest/agile/1.0/sprint/%d", j.domain, id)
	sprint := &Sprint{}
	res, err := j.request("GET", uri, nil, sprint)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status was not OK: %s\n", res.Status))
	}

	return sprint, nil
}
