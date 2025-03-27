package models

type User struct {
	Username string
	Password string
}

type UserWithID struct {
	ID       string
	Username string
}

type UserWithPassword struct {
	ID       string
	Username string
	Password []byte
}
