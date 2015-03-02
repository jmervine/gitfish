package main

import (
	"net/http"
)

type Conditions struct {
	Branches []string
	Token    string
	Owner    bool
	Admin    bool
	Master   bool
}

func (c Conditions) Auth(r *http.Request) bool {
	// TODO: check and see if github support basic auth and if
	// so replace simple token auth with basic auth.

	if c.Token == "" {
		return true
	}

	// intentionally restrictive
	var token string
	for k := range r.URL.Query() {
		token = k
		break // get first key
	}

	if c.Token == token {
		return true
	}

	return false
}

func (c Conditions) AreMet(p PushEvent) bool {
	found := true
	if len(c.Branches) != 0 {
		found = false
		for _, branch := range c.Branches {
			if p.Branch() == branch {
				found = true
				break
			}
		}
	}

	if !found {
		return false
	}

	if c.Owner && !p.ByOwner() {
		return false
	}

	if c.Admin && !p.ByAdmin() {
		return false
	}

	if c.Master && !p.ToMaster() {
		return false
	}

	return true
}
