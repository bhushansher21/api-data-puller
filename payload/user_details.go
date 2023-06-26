package payload

// UserList struct is storing user details in json file.
type UserList []struct {
	ID                 int    `json:"id"`
	Nickname           string `json:"nickname"`
	GravatarID         string `json:"gravatar_id"`
	GithubProfile      string `json:"github_profile"`
	TwitterProfile     any    `json:"twitter_profile"`
	ContributionsCount int    `json:"contributions_count"`
	Organisations      `json:"organisations"`
	Link               string `json:"link"`
	PullRequests       `json:"pull_requests"`
}

type Organisations []struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Link      string `json:"link"`
}

type PullRequests []struct {
	Title     string `json:"title"`
	IssueURL  string `json:"issue_url"`
	RepoName  string `json:"repo_name"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}
