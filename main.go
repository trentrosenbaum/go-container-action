package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"os"
)

func main() {
	myInput := os.Getenv("INPUT_MYINPUT")

	output := fmt.Sprintf("Hello %s", myInput)

	fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))

	client := github.NewClient(nil)

	reviewers, _, err := client.PullRequests.ListReviewers(context.Background(), "trentrosenbaum", "go-container-action", 1, &github.ListOptions{})
	if err != nil {
		return
	}

	teams := reviewers.Teams
	users := reviewers.Users

	for i, v := range teams {
		fmt.Printf("Team %d: %s", i, v.Name)
	}

	for i, v := range users {
		fmt.Printf("User %d: %s", i, v.Name)
	}

}
