package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Gitubrr/GoSymGym/internal/client"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

func printUsage() {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "GitHub Repository Info CLI\n\n")
	fmt.Fprintf(out, "usage: GoSymGym [-h] -o REPO_OWNER -n REPO_NAME [-t GITHUB_TOKEN] [-T TIMEOUT]\n\n")
	fmt.Fprintf(out, "options:\n")
	fmt.Fprintf(out, "\t-h, --help\t\tshow this help message and exit\n")
	fmt.Fprintf(out, "\t-o REPO_OWNER\t\tName of the repository owner\n")
	fmt.Fprintf(out, "\t-n REPO_NAME\t\tRepository name\n")
	fmt.Fprintf(out, "\t-t GITHUB_TOKEN\t\tPersonal access token\n")
	fmt.Fprintf(out, "\t-T TIMEOUT\t\tMaximum response time sec\n")
}

func main() {
	ownerFlag := flag.String("o", "", "")
	nameFlag := flag.String("n", "", "")
	tokenFlag := flag.String("t", "", "")
	timeoutFlag := flag.Int("T", 10, "")

	flag.Usage = printUsage

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

	client := client.NewClient(token, timeout)

	info, err := client.GetRepoInfo(repoOwner, repoName)
	if err != nil {
		fmt.Fprintf(os.Stderr, red+"Error: %v\n"+reset, err)
		os.Exit(1)
	}

	fmt.Println(info)
}
