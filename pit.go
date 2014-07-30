package main

import (
	"github.com/typester/go-pit"
)

func pitToken(user string) string {
	if user == "" {
		return ""
	}
	profile, err := pit.Get("github-release."+user, pit.Requires{"token": "github access token for github-release"})
	if err != nil {
		return ""
	}
	return (*profile)["token"]
}
