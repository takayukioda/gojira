package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Jira struct {
	client *http.Client
	domain string
	email  string
	token  string
}

func NewJira(domain, email, token string) *Jira {
	return &Jira{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		domain: domain,
		email:  email,
		token:  token,
	}
}

func (j *Jira) request(method, uri string, body io.Reader, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header = map[string][]string{
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
	}
	req.SetBasicAuth(j.email, j.token)

	res, err := j.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if v == nil {
		return res, nil
	}

	if w, ok := v.(io.Writer); ok {
		// Give *bytes.Buffer to get raw bytes instead of decoded struct
		io.Copy(w, res.Body)
	} else {
		err := json.NewDecoder(res.Body).Decode(v)

		if err != nil {
			return res, nil
		}
	}
	return res, nil
}
