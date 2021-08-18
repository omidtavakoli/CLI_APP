package integration

import (
	"CLI_APP/model"
	"sync"
)

func IntegrateDate(events []*model.Event, commits []*model.Commit, repos []*model.Repo, actors []*model.Actor) {
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
				for _, repo := range repos {
					if repo.ID == event.RepoID {
						event.Repo = repo
					}
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
