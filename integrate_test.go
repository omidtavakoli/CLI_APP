package main

import "testing"

func TestRepoExistence(t *testing.T) {
	repos, commits, actors, events, err := FetchData()
	if err != nil {
		panic(err.Error())
	}

	CleanData(&events, &commits, &repos, &actors)
	IntegrateDate(events, commits, repos, actors)

	for _, event := range events {
		if event.Repo == nil {
			t.Errorf("eventID:%d repo is nil", event.ID)
			break
		}
	}
}
