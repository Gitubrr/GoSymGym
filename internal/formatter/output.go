package formatter

import (
	"fmt"

	"github.com/Gitubrr/GoSymGym/internal/client"
)

func PrintRepoInfo(info *client.RepoInfo) {

	fmt.Printf("Repository name:\t%s\n", info.RepoName)
	if info.Description != "" {
		fmt.Printf("Description:\t\t%s\n", info.Description)
	} else {
		fmt.Printf("Description:\t\tThis repository has no description\n")
	}
	fmt.Printf("Stars:\t\t\t%d\n", info.Stars)
	fmt.Printf("Forks:\t\t\t%d\n", info.Forks)
	fmt.Printf("Issues:\t\t\t%d\n", info.Issues)
	if info.Language != "" {
		fmt.Printf("Language:\t\t%s\n", info.Language)
	}
	fmt.Printf("Created:\t\t%s\n", info.CreatedAt.Format("2006-01-02"))
	fmt.Printf("Updated:\t\t%s\n", info.UpdatedAt.Format("2006-01-02"))
	fmt.Printf("URL:\t\t\t%s\n", info.HTMLURL)
}
