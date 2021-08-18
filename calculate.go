package main

import (
	"sort"
	"sync"
)

func topActiveUsers(users []*Actor, count int) []string {
	var actors []Actor
	for i, user := range users {
		for _, event := range user.Events {
			if event.Type == "PushEvent" || event.Type == "PullRequestEvent" {
				users[i].Activity += 1
			}
		}
		if user.Activity > 0 {
			actors = append(actors, Actor{
				ID:       user.ID,
				UserName: user.UserName,
				Events:   nil,
				Activity: user.Activity,
			})
		}
	}
	sort.Slice(actors, func(i, j int) bool {
		return actors[i].Activity > actors[j].Activity
	})

	var topUsers []string
	for _, user := range actors[:count-1] {
		topUsers = append(topUsers, user.UserName)
	}

	return topUsers
}

func topActiveRepos(events []*Event, repos []*Repo, count int) ([]string, []string) {
	var reposByCommit []Repo

	for _, event := range events {
		for j, repo := range repos {
			if event.Repo != nil {
				if event.Repo.ID == repo.ID {
					if event.Type == "CommitCommentEvent" {
						repos[j].CommitsCount += 1
					} else if event.Type == "WatchEvent" {
						repos[j].WatchesCount += 1
					}
				}
			}
		}
	}

	for _, repo := range repos {
		if repo.CommitsCount > 0 || repo.WatchesCount > 0 {
			reposByCommit = append(reposByCommit, Repo{
				ID:           repo.ID,
				Name:         repo.Name,
				CommitsCount: repo.CommitsCount,
				WatchesCount: repo.WatchesCount,
			})
		}
	}

	reposByWatches := make([]Repo, len(reposByCommit))
	copy(reposByWatches, reposByCommit)

	sort.Slice(reposByWatches, func(i, j int) bool {
		return reposByWatches[i].WatchesCount > reposByWatches[j].WatchesCount
	})
	sort.Slice(reposByCommit, func(i, j int) bool {
		return reposByCommit[i].CommitsCount > reposByCommit[j].CommitsCount
	})

	var rbc []string
	for _, repo := range reposByCommit[:count-1] {
		rbc = append(rbc, repo.Name)
	}

	var rbw []string
	for _, repo := range reposByWatches[:count-1] {
		rbw = append(rbw, repo.Name)
	}
	return rbc, rbw
}

func Calculation(repos []*Repo, actors []*Actor, events []*Event, count int) Stats {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var users []string
	go func() {
		defer wg.Done()
		users = topActiveUsers(actors, count)
	}()

	var topReposByCommit, topReposByWatches []string
	go func() {
		defer wg.Done()
		topReposByCommit, topReposByWatches = topActiveRepos(events, repos, count)
	}()

	wg.Wait()

	return Stats{
		TopUsers:          users,
		TopReposByCommit:  topReposByCommit,
		TopReposByWatches: topReposByWatches,
	}
}
