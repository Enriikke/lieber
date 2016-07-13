package marvel

import (
	"encoding/json"
	"fmt"
)

const storyPath = "/v1/public/stories"

// Story represents an indivisible, reusable components of comics.
type Story struct {
}

// StoryList uses the Marvel API to return a slice of Story structs.
func (client Client) StoryList() ([]Story, error) {
	resp, err := client.sendRequest("GET", storyPath)
	if err != nil {
		return nil, err
	}

	var stories []Story
	err := json.Unmarshal(resp, &stories)
	if err != nil {
		return nil, err
	}

	return stories, nil
}
