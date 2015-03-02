package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"strings"
)

type Conditions struct {
	Branches []string
	Secret   string
	Owner    bool
	Admin    bool
	Master   bool
}

func (c Conditions) Auth(body []byte, sig string) bool {
	if c.Secret == "" {
		return true
	}

	if sig == "" {
		return false
	}

	hasher := hmac.New(sha1.New, []byte(c.Secret))
	hasher.Write(body)
	expected := []byte(hex.EncodeToString(hasher.Sum(nil)))
	recieved := []byte(strings.Split(sig, "=")[1])

	passed := hmac.Equal(expected, recieved)

	if !passed {
		log.Println("X-Hub-Signature check failed")
	}
	return passed
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
