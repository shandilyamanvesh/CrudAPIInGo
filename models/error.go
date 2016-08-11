package models

type Error struct {
	Code string
	Msg  string
}

func (err *Error) Error() string {
	return err.Msg
}
