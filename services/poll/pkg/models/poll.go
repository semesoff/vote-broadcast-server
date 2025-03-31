package models

type PollType int

const (
	single PollType = iota
	multiple
)

func (p PollType) String() string {
	switch p {
	case single:
		return "single"
	case multiple:
		return "multiple"
	default:
		return "unknown"
	}
}

const (
	MaxPollOptions = 10
	MaxPollType    = 1
	MinPollType    = 0
	MaxPollTitle   = 100
	MinPollTitle   = 1
	MaxOptionText  = 100
	MinOptionText  = 1
)

type Poll struct {
	ID         int
	Title      string
	Type       PollType
	MaxOptions int
	Options    []Option
}

type Option struct {
	ID   int
	Text string
}
