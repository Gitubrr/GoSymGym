package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Gitubrr/GoSymGym/internal/client"
	"github.com/Gitubrr/GoSymGym/internal/formatter"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

func main() {
	ownerFlag := flag.String("o", "", "")
	nameFlag := flag.String("n", "", "")
	tokenFlag := flag.String("t", "", "")
	timeoutFlag := flag.Int("T", 10, "")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "GitHub Repository Info CLI\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "usage: GoSymGym [-h] -o REPO_OWNER -n REPO_NAME [-t GITHUB_TOKEN] [-T TIMEOUT]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "options:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t-h, --help\t\tshow this help message and exit\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t-o REPO_OWNER\t\tName of the repository owner\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t-n REPO_NAME\t\tRepository name\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t-t GITHUB_TOKEN\t\tPersonal access token\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t-T TIMEOUT\t\tMaximum response time sec\n")
	}

	flag.Parse()

	var repoOwner, repoName, token string
	var timeout int

	repoOwner = *ownerFlag
	repoName = *nameFlag
	token = *tokenFlag
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	timeout = *timeoutFlag

	client := client.NewClient()
	if token != "" {
		client.SetToken(token)
	}
	if timeout != 10 {
		client.SetTimeout(time.Duration(timeout) * time.Second)
	}

	info, err := client.GetRepoInfo(repoOwner, repoName)
	if err != nil {
		fmt.Fprintf(os.Stderr, red+"Error: %v\n"+reset, err)
		os.Exit(1)
	}

	formatter.PrintRepoInfo(info)
}
