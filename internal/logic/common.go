package logic

// Error represents logic error, user must do something.
type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}
