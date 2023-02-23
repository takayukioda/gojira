package main

import (
	"errors"
	"fmt"
)

type Board struct {
	Id int `json:"id"`
}

// Get a board with a given ID. See [ref: get a board] for more details.
//
// [ref: get a board]: https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-rest-agile-1-0-board-boardid-get
func (j *Jira) GetBoard(id int64) (*Board, error) {
	uri := fmt.Sprintf("https://%s/rest/agile/1.0/board/%d", j.domain, id)
	board := &Board{}
	res, err := j.request("GET", uri, nil, board)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status was not OK: %s\n", res.Status))
	}

	return board, nil
}
