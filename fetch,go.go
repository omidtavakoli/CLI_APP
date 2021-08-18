package main

import (
	"github.com/gocarina/gocsv"
	"os"
)

func loadCSV(filePath string, out interface{}) error {
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

func FetchData() ([]*Repo, []*Commit, []*Actor, []*Event, error) {
	var repos []*Repo
	err := loadCSV("./data/repos.csv", &repos)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var commits []*Commit
	err = loadCSV("./data/commits.csv", &commits)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var actors []*Actor
	err = loadCSV("./data/actors.csv", &actors)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var events []*Event
	err = loadCSV("./data/events.csv", &events)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return repos, commits, actors, events, nil
}
