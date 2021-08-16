package main

type User struct {
	ID       int64
	UserName string
}

type Commit struct {
	SHA     string
	Message string
	EventID int64
}

type Events struct {
	ID     int64
	Type   string
	UserID int64
	RepoID int64
}

type Repo struct {
	ID   int64
	Name string
}
