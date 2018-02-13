package jira

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

const BaseUrl = "http://jirainstance"

// Send request to Jira. jsonStr represents the body of the request and is the JiraLog struct.
func SendRequest(jsonStr []byte, ticket string, username string, password string) *http.Response {
	req, err := http.NewRequest("POST", BuildUrlString(ticket), bytes.NewBuffer(jsonStr))
	CheckErr(err)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	CheckResponseStatus(resp, ticket)

	return resp
}

func BuildUrlString(ticket string) string {
	urlComponents := []string{BaseUrl, ticket, "/worklog"}

	return strings.Join(urlComponents, "")
}

func CheckResponseStatus(resp *http.Response, ticket string) {
	if resp.StatusCode >= 300 {
		fmt.Println(fmt.Sprintf("Something went wrong with the Jira request for: %s", ticket))
	}

	if resp.StatusCode == 401 {
		fmt.Println("Invalid credentials provided")
		os.Exit(1)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}