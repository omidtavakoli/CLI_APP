package main

import (
	"github.com/gocarina/gocsv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"log"
	"os"
	"sync"
)

var app = cli.NewApp()

func info() {
	app.Name = "Users/Repositories CLI App"
	app.Usage = "A simple CLI for getting users'/repositories' stats"
	app.Author = "Omid Tavakoli"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "top-10-users",
			Aliases: []string{"tu"},
			Usage:   "Top 10 active users sorted by amount of PRs created and commits pushed",
			Action: func(c *cli.Context) {
				//todo: implement me
				logrus.Info("calculate top-10-users here ...")
			},
		},
		{
			Name:    "top-10-repositories-by-commit",
			Aliases: []string{"tr"},
			Usage:   "Top 10 repositories sorted by amount of commits pushed",
			Action: func(c *cli.Context) {
				//todo: implement me
				logrus.Info("calculate top-10-repositories-by-commit here ...")
			},
		},
		{
			Name:    "top-10-repositories-by-watch",
			Aliases: []string{"trw"},
			Usage:   "Top 10 repositories sorted by amount of watch events",
			Action: func(c *cli.Context) {
				//todo: implement me
				logrus.Info("calculate top-10-repositories-by-watch here ...")
			},
		},
	}
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

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

func fetchData() ([]*Repo, []*Commit, []*Actor, []*Event, error) {
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

func cleanData(events []*Event, commits []*Commit, repos []*Repo, actors []*Actor) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, event := range events {
			wg2 := sync.WaitGroup{}
			wg2.Add(2)
			go func() {
				defer wg2.Done()
				var commitsList []string
				var uniqueCommits []*Commit
				for _, commit := range commits {
					_, find := Find(commitsList, commit.SHA)
					if !find {
						if commit.EventID == event.ID {
							event.Message = commit.Message
							event.SHA = commit.SHA
						}
						commitsList = append(commitsList, commit.SHA)
						uniqueCommits = append(uniqueCommits, commit)
					}
				}
				commits = uniqueCommits
			}()
			go func() {
				defer wg2.Done()
				var reposList []string
				var uniqueRepos []*Repo
				for _, repo := range repos {
					_, find := Find(reposList, repo.Name)
					if !find {
						if repo.ID == event.RepoID {
							event.Repo = repo
						}
						reposList = append(reposList, repo.Name)
						uniqueRepos = append(uniqueRepos, repo)
					}
				}
			}()
			wg2.Wait()
		}
	}()

	go func() {
		defer wg.Done()

		//used for ignoring duplicates
		var usersList []string
		var uniqueActors []*Actor
		for _, user := range actors {
			_, find := Find(usersList, user.UserName)
			if !find {
				for _, event := range events {
					if event.ActorID == user.ID {
						user.Events = append(user.Events, event)
					}
				}
				usersList = append(usersList, user.UserName)
				uniqueActors = append(uniqueActors, user)
			}
		}
		actors = uniqueActors
	}()

	wg.Wait()
}

func main() {
	info()
	repos, commits, actors, events, err := fetchData()
	if err != nil {
		panic(err.Error())
	}

	cleanData(events, commits, repos, actors)

	logrus.Info("len: ", len(repos), len(commits), len(actors), len(events))

	commands()
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
