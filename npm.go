package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

type npmQuery struct {
	Pkg string `form:"pkg" json:"pkg"`
}

type npmPkg struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DistTags    struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

func (q *npmQuery) check() string {
	version := "unknown"
	url := "https://registry.npmjs.com/" + q.Pkg
	resp, err := resty.New().R().Get(url)
	if err != nil {
		log.Printf("request error: %v\n", err)
		return version
	}
	body := resp.Body()
	var pkg npmPkg
	if err := json.Unmarshal(body, &pkg); err != nil {
		log.Printf("response error: %v\n", err)
		return version
	}
	log.Printf("response: %v\n", pkg)
	if pkg.DistTags.Latest != "" {
		version = pkg.DistTags.Latest
	}
	return version
}
