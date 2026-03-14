package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrNotFound     = fmt.Errorf("repository not found")
	ErrRateLimit    = fmt.Errorf("rate limit exceeded")
	ErrUnauthorized = fmt.Errorf("unauthorized - check your token")
)

type RepoInfo struct {
	RepoName    string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	Issues      int       `json:"open_issues_count"`
	Language    string    `json:"language"`
	HTMLURL     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	userAgent  string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   "https://api.github.com",
		userAgent: "GoSymGym-CLI/1.0",
	}
}

func (client *Client) GetRepoInfo(owner, repo string) (*RepoInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", client.baseURL, owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("User-Agent", client.userAgent)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	if client.token != "" {
		req.Header.Set("Authorization", "Bearer "+client.token)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var info RepoInfo
		if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
			return nil, fmt.Errorf("parsing JSON: %w", err)
		}
		return &info, nil
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusForbidden:
		return nil, ErrRateLimit
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	default:
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
}

func (client *Client) SetToken(token string) {
	client.token = token
}

func (client *Client) SetTimeout(timeout time.Duration) {
	client.httpClient.Timeout = timeout
}
