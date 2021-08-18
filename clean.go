package main

import (
	"sync"
)

func CleanData(events *[]*Event, commits *[]*Commit, repos *[]*Repo, actors *[]*Actor) {
	wg := sync.WaitGroup{}
	wg.Add(4)

	//unify actors
	go func() {
		defer wg.Done()
		*actors = CleanActors(*actors)
	}()

	//unify events
	go func() {
		defer wg.Done()
		*events = CleanEvents(*events)
	}()

	//unify commits
	go func() {
		defer wg.Done()
		*commits = CleanCommits(*commits)
	}()

	//unify repos
	go func() {
		defer wg.Done()
		*repos = CleanRepos(*repos)
	}()

	wg.Wait()
}

func CleanActors(actors []*Actor) []*Actor {
	var uniqueActors []*Actor
	actorsMap := make(map[int64]*Actor)

	for _, user := range actors {
		if _, ok := actorsMap[user.ID]; !ok {
			uniqueActors = append(uniqueActors, user)
			actorsMap[user.ID] = user
		}
	}

	return uniqueActors
}

func CleanRepos(repos []*Repo) []*Repo {
	var uniqueRepos []*Repo
	//just another type of implementation
	var reposList IntegerSlice

	for _, repo := range repos {
		_, found := reposList.Find(repo.ID)
		if !found {
			uniqueRepos = append(uniqueRepos, repo)
			reposList.slice = append(reposList.slice, repo.ID)
		}
	}

	return uniqueRepos
}

func CleanEvents(events []*Event) []*Event {
	var uniqueEvents []*Event
	eventsMap := make(map[int64]*Event)

	for _, event := range events {
		if _, ok := eventsMap[event.ID]; !ok {
			uniqueEvents = append(uniqueEvents, event)
			eventsMap[event.ID] = event
		}
	}

	return uniqueEvents
}

func CleanCommits(commits []*Commit) []*Commit {
	var uniqueCommits []*Commit
	commitsMap := make(map[int64]*Commit)

	for _, commit := range commits {
		if _, ok := commitsMap[commit.EventID]; !ok {
			uniqueCommits = append(uniqueCommits, commit)
			commitsMap[commit.EventID] = commit
		}
	}

	return uniqueCommits
}
