package models

type Err struct {
	Message string
}
type ErrEmptyRequest struct {
	Err
}

type ErrEmptyData struct {
	Err
}

type ErrInvalidData struct {
	Err
}

func (e Err) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "data is invalid"
}

func (e ErrEmptyData) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "data cannot be empty"
}

func (e ErrEmptyRequest) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "request cannot be nil"
}
