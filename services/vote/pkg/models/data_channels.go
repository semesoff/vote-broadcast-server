package models

const (
	GetVotes MethodType = iota
)

type MethodType int

type DataChannel chan Data

func (methodType MethodType) String() string {
	switch methodType {
	case GetVotes:
		return "getVotes"
	default:
		return "unknown"
	}
}

type Data struct {
	Method MethodType
	Data   interface{}
}
