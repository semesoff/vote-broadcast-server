package models

type PollVotes struct {
	ID      int
	Options map[int]Option
}

type Option struct {
	CountVotes int
	Users      []User
}

type User struct {
	ID   int
	Name string
}
