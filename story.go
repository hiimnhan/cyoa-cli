package main

import (
	"encoding/json"
	"fmt"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func parseStory(storyMap []byte) (story Story) {
	if err := json.Unmarshal(storyMap, &story); err != nil {
		fmt.Printf("Error parsing story map %s", err)
	}
	return
}
