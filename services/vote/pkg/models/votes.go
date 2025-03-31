package models

type PollType int

const (
	Single PollType = iota
	Multiple
)

func (p PollType) String() string {
	switch p {
	case Single:
		return "single"
	case Multiple:
		return "multiple"
	default:
		return "unknown"
	}
}

type Votes map[int]Vote

type Vote struct {
	OptionId   int
	CountVotes int
	Users      []User
}

type User struct {
	ID   int
	Name string
}

type UserVote struct {
	PollId    int
	UserId    int
	OptionsId []int
}
