package main

import "strings"

// Note: comment out that which I don't need currently.
// For json see: https://developer.github.com/v3/activity/events/types/#pushevent
type PushEvent struct {
	Commits []struct {
		Added     []string `json:"added"`
		Modified  []string `json:"modified"`
		Removed   []string `json:"removed"`
		Timestamp string   `json:"timestamp"`
	} `json:"commits"`
	HeadCommit struct {
		Added     []string `json:"added"`
		Modified  []string `json:"modified"`
		Removed   []string `json:"removed"`
		Timestamp string   `json:"timestamp"`
	} `json:"head_commit"`
	Sender struct {
		Login     string `json:"login"`
		SiteAdmin bool   `json:"site_admin"`
		Type      string `json:"type"`
	} `json:"sender"`
	Ref        string `json:"ref"`
	Repository struct {
		Name         string `json:"name"`
		MasterBranch string `json:"master_branch"`
		Owner        struct {
			Name string `json:"name"`
		} `json:"owner"`
	} `json:"repository"`
}

func (e PushEvent) Branch() (b string) {
	prefix := "refs/heads/"
	if strings.HasPrefix(e.Ref, prefix) {
		b = strings.TrimPrefix(e.Ref, prefix)
	}

	return
}

func (e PushEvent) ByOwner() bool {
	return e.Sender.Login == e.Repository.Owner.Name
}

func (e PushEvent) ByAdmin() bool {
	return e.Sender.SiteAdmin
}

func (e PushEvent) ToMaster() bool {
	return e.Branch() == e.Repository.MasterBranch
}
