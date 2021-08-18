package main

import (
	"github.com/sirupsen/logrus"
	"sync"
)

func IntegrateDate(events []*Event, commits []*Commit, repos []*Repo, actors []*Actor) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, event := range events {
			wg2 := sync.WaitGroup{}
			wg2.Add(3)

			go func() {
				defer wg2.Done()
				for _, commit := range commits {
					if commit.EventID == event.ID {
						event.Message = commit.Message
						event.SHA = commit.SHA
					}
				}
			}()

			go func() {
				defer wg2.Done()
				for _, repo := range repos {
					if repo.ID == event.RepoID {
						event.Repo = repo
					}
				}
			}()

			go func() {
				defer wg2.Done()
				done := false
				for _, repo := range repos {
					if repo.ID == event.RepoID {
						done = true
						event.Repo = repo
					}
				}
				if !done {
					logrus.Info(event)
				}
			}()

			wg2.Wait()
		}
	}()

	go func() {
		defer wg.Done()
		for _, user := range actors {
			for _, event := range events {
				if event.ActorID == user.ID {
					user.Events = append(user.Events, event)
				}
			}
		}
	}()

	wg.Wait()
}
