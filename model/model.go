package model

type Actor struct {
	ID       int64  `csv:"id"`
	UserName string `csv:"username"`
	Events   []*Event
	Activity int
}

// Event : consider commit as an event
type Event struct {
	ID      int64  `csv:"id"`
	Type    string `csv:"type"`
	ActorID int64  `csv:"actor_id"`
	RepoID  int64  `csv:"repo_id"`
	SHA     string
	Message string
	Repo    *Repo
}

type Repo struct {
	ID           int64  `csv:"id"`
	Name         string `csv:"name"`
	CommitsCount int64
	WatchesCount int64
}

type Commit struct {
	SHA     string `csv:"sha"`
	Message string `csv:"message"`
	EventID int64  `csv:"event_id"`
}

type Stats struct {
	TopUsers          []string
	TopReposByCommit  []string
	TopReposByWatches []string
}
