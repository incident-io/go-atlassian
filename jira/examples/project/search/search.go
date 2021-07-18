package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	options := &jira.ProjectSearchOptionsScheme{
		OrderBy: "issueCount",
		Action:  "browse",
		Expand:  []string{"insight", "lead", "issueTypes", "projectKeys", "description"},
	}

	var (
		startAt    = 0
		maxResults = 50
	)

	projects, response, err := atlassian.Project.Search(context.Background(), options, startAt, maxResults)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, project := range projects.Values {
		log.Println(project)
	}
}
