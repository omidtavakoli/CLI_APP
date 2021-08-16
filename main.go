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

func Calculation() error {
	//todo: load data to memory
	logrus.Warn("implement me : calculation")
	return nil
}

func main() {
	info()
	commands()
	err := Calculation()
	if err != nil {
		logrus.Error(err)
		panic("loading data error")
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
