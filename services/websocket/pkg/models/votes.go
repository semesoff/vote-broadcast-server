package models

type Option struct {
	ID         int
	CountVotes int
	Users      []User
}

type User struct {
	ID   int
	Name string
}
