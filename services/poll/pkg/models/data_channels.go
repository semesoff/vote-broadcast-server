package models

const (
	GetPolls MethodType = iota
)

type MethodType int

type DataChannel chan Data

func (methodType MethodType) String() string {
	switch methodType {
	case GetPolls:
		return "getPolls"
	default:
		return "unknown"
	}
}

type Data struct {
	Method MethodType
	Data   interface{}
}
