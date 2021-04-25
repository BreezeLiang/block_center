package common


type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}

