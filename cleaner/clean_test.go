package cleaner

import (
	"CLI_APP/fetch"
	"errors"
	"fmt"
	"testing"
)

func TestCleanRepos(t *testing.T) {
	repos, _, _, _, err := fetch.FetchData()
	if err != nil {
		panic(err.Error())
	}
	cleanRepos := CleanRepos(repos)
	reposMap := make(map[int64]string)
	for _, repo := range cleanRepos {
		reposMap[repo.ID] = repo.Name
	}
	var errs []error
	for _, repo := range repos {
		if _, ok := reposMap[repo.ID]; !ok {
			msg := fmt.Sprintf("repo:%d not exist in cleansRepo", repo.ID)
			errs = append(errs, errors.New(msg))
		}
	}
	if len(errs) > 0 {
		t.Errorf("some repos are not in the clean one")
	}
}

func TestCleanActors(t *testing.T) {
	_, _, actors, _, err := fetch.FetchData()
	if err != nil {
		panic(err.Error())
	}
	cleanRepos := CleanActors(actors)
	actorsMap := make(map[int64]int64)
	for _, actor := range cleanRepos {
		actorsMap[actor.ID] = actor.ID
	}
	var errs []error
	for _, actor := range actors {
		if _, ok := actorsMap[actor.ID]; !ok {
			msg := fmt.Sprintf("actor:%d not exist in cleansRepo", actor.ID)
			errs = append(errs, errors.New(msg))
		}
	}
	if len(errs) > 0 {
		t.Errorf("some actors are not in the clean one")
	}
}

func TestCleanCommits(t *testing.T) {
	_, commits, _, _, err := fetch.FetchData()
	if err != nil {
		panic(err.Error())
	}
	cleanRepos := CleanCommits(commits)
	commitsMap := make(map[int64]string)
	for _, actor := range cleanRepos {
		commitsMap[actor.EventID] = actor.SHA
	}
	var errs []error
	for _, commit := range commits {
		if _, ok := commitsMap[commit.EventID]; !ok {
			msg := fmt.Sprintf("commit:%d not exist in cleansRepo", commit.EventID)
			errs = append(errs, errors.New(msg))
		}
	}
	if len(errs) > 0 {
		t.Errorf("some commits are not in the clean one")
	}
}

func TestCleanEvents(t *testing.T) {
	_, _, _, events, err := fetch.FetchData()
	if err != nil {
		panic(err.Error())
	}
	cleanRepos := CleanEvents(events)
	eventsMap := make(map[int64]string)
	for _, event := range cleanRepos {
		eventsMap[event.ID] = event.Type
	}
	var errs []error
	for _, event := range events {
		if _, ok := eventsMap[event.ID]; !ok {
			msg := fmt.Sprintf("event:%d not exist in cleansRepo", event.ID)
			errs = append(errs, errors.New(msg))
		}
	}
	if len(errs) > 0 {
		t.Errorf("some events are not in the clean one")
	}
}

func TestReposExistence(t *testing.T) {
	repos, _, _, _, err := fetch.FetchData()
	if err != nil {
		panic(err.Error())
	}
	cleanRepos := CleanRepos(repos)

	tests := []struct {
		description string
		name        string
		want        int64
	}{
		{
			description: "repo existence",
			name:        "success",
			want:        83585690,
		},
		{
			description: "repo existence",
			name:        "success",
			want:        231161702,
		},
		{
			description: "repo existence",
			name:        "success",
			want:        227863306,
		},
		{
			description: "repo existence",
			name:        "success",
			want:        86929735,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			found := false
			for _, repo := range cleanRepos {
				if repo.ID == tt.want {
					found = true
				}
			}
			if !found {
				t.Errorf("repo:%d not found in repos", tt.want)
			}
		})
	}
}
