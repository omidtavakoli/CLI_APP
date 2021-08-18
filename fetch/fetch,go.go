package fetch

import (
	"CLI_APP/model"
	_ "embed"
	"github.com/gocarina/gocsv"
	"os"
)

func LoadCSV(filePath string, out interface{}) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, out); err != nil {
		return err
	}

	return nil
}

func FetchData() ([]*model.Repo, []*model.Commit, []*model.Actor, []*model.Event, error) {
	var repos []*model.Repo
	err := LoadCSV("./data/repos.csv", &repos)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var commits []*model.Commit
	err = LoadCSV("./data/commits.csv", &commits)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var actors []*model.Actor
	err = LoadCSV("./data/actors.csv", &actors)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var events []*model.Event
	err = LoadCSV("./data/events.csv", &events)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return repos, commits, actors, events, nil
}
