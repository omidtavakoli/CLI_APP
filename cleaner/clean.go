package cleaner

import (
	"CLI_APP/general"
	"CLI_APP/model"
	"sync"
)

func CleanData(events *[]*model.Event, commits *[]*model.Commit, repos *[]*model.Repo, actors *[]*model.Actor) {
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

func CleanActors(actors []*model.Actor) []*model.Actor {
	var uniqueActors []*model.Actor
	actorsMap := make(map[int64]*model.Actor)

	for _, user := range actors {
		if _, ok := actorsMap[user.ID]; !ok {
			uniqueActors = append(uniqueActors, user)
			actorsMap[user.ID] = user
		}
	}

	return uniqueActors
}

func CleanRepos(repos []*model.Repo) []*model.Repo {
	var uniqueRepos []*model.Repo
	//just another type of implementation
	var reposList general.IntegerSlice

	for _, repo := range repos {
		_, found := reposList.Find(repo.ID)
		if !found {
			uniqueRepos = append(uniqueRepos, repo)
			reposList.Slice = append(reposList.Slice, repo.ID)
		}
	}

	return uniqueRepos
}

func CleanEvents(events []*model.Event) []*model.Event {
	var uniqueEvents []*model.Event
	eventsMap := make(map[int64]*model.Event)

	for _, event := range events {
		if _, ok := eventsMap[event.ID]; !ok {
			uniqueEvents = append(uniqueEvents, event)
			eventsMap[event.ID] = event
		}
	}

	return uniqueEvents
}

func CleanCommits(commits []*model.Commit) []*model.Commit {
	var uniqueCommits []*model.Commit
	commitsMap := make(map[int64]*model.Commit)

	for _, commit := range commits {
		if _, ok := commitsMap[commit.EventID]; !ok {
			uniqueCommits = append(uniqueCommits, commit)
			commitsMap[commit.EventID] = commit
		}
	}

	return uniqueCommits
}
