package models

type Vote struct {
	PollId    int   `json:"poll_id"`
	UserId    int   `json:"user_id"`
	OptionsId []int `json:"options_id"`
}

type VoteRequest struct {
	PollId    int    `json:"poll_id"`
	OptionsId []int  `json:"options_id"`
	Token     string `json:"token"`
}
