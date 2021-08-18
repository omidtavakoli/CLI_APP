package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Users/Repositories CLI App"
	app.Usage = "A simple CLI for getting users'/repositories' stats"
	app.Author = "Omid Tavakoli"
	app.Version = "1.0.0"
}

func commands(stats Stats) {
	app.Commands = []cli.Command{
		{
			Name:    "top-10-users",
			Aliases: []string{"tu"},
			Usage:   "Top 10 active users sorted by amount of PRs created and commits pushed",
			Action: func(c *cli.Context) {
				logrus.Infof("Top users by by amount of PRs created and commits pushed are : %s", CleanText(stats.TopUsers))
			},
		},
		{
			Name:    "top-10-repositories-by-commit",
			Aliases: []string{"tr"},
			Usage:   "Top 10 repositories sorted by amount of commits pushed",
			Action: func(c *cli.Context) {
				logrus.Infof("Top repositories sorted by amount of commits pushed are : %s", CleanText(stats.TopReposByCommit))
			},
		},
		{
			Name:    "top-10-repositories-by-watch",
			Aliases: []string{"trw"},
			Usage:   "Top 10 repositories sorted by amount of watch events",
			Action: func(c *cli.Context) {
				logrus.Infof("Top repositories sorted by amount of watch events are : %s", CleanText(stats.TopReposByWatches))
			},
		},
	}
}

func main() {
	//showing cli info
	info()

	//loading data to memory
	repos, commits, actors, events, err := FetchData()
	if err != nil {
		panic(err.Error())
	}

	//cleaning data
	CleanData(&events, &commits, &repos, &actors)

	//enrich data
	IntegrateDate(events, commits, repos, actors)

	//calculate commands
	stats := Calculation(repos, actors, events, 10)

	//prepare commands
	logrus.Infof("data prepared => events:%d, commits:%d, repos:%d, actors:%d", len(events), len(commits), len(repos), len(actors))
	commands(stats)

	//run cli app
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
