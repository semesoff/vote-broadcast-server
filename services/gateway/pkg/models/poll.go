package models

type PollType int

const (
	single PollType = iota
	multiple
)

type Poll struct {
	Title   string   `json:"title"`
	Type    PollType `json:"type"`
	Options []Option `json:"options"`
	UserID  int      `json:"user_id"`
}

type Option struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type PollRequest struct {
	Title   string   `json:"title"`
	Type    PollType `json:"type"`
	Options []Option `json:"options"`
}
