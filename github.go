package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
	"log"
)

type gitHubQuery struct {
	Repo       string `form:"repo" json:"repo"`
	Prerelease bool   `form:"prerelease" json:"prerelease"`
}

type gitHubRelease struct {
	TagName    string `json:"tag_name"`
	Prerelease bool   `json:"prerelease"`
}

func (q *gitHubQuery) check() string {
	version := "unknown"
	url := "https://api.github.com/repos/" + q.Repo + "/releases"
	resp, err := resty.New().R().Get(url)
	if err != nil {
		log.Printf("request error: %v\n", err)
		return version
	}
	body := resp.Body()
	releases := make([]gitHubRelease, 0)
	if err := json.Unmarshal(body, &releases); err != nil {
		log.Printf("response error: %v", err)
		return version
	}
	log.Printf("response: %v\n", string(body))
	// include prerelease
	if q.Prerelease {
		latest, find := lo.Find(releases, func(r gitHubRelease) bool {
			return true
		})
		if find {
			return latest.TagName
		}
	}
	latest, find := lo.Find(releases, func(r gitHubRelease) bool {
		return !r.Prerelease
	})
	if find {
		version = latest.TagName
	}
	return version
}
