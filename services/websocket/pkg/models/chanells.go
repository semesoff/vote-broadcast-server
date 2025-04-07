package models

const (
	GetPolls MethodType = iota
	GetVotes
)

type MethodType int

func (methodType MethodType) String() string {
	switch methodType {
	case GetPolls:
		return "getPolls"
	case GetVotes:
		return "getVotes"
	default:
		return "unknown"
	}
}

type DataChannels chan Data

type Data struct {
	Method MethodType
	Data   interface{}
}
