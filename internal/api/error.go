package api

type Error struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

func (e Error) Error() string { return e.Err }

func NewError(msg string, err error) Error {
	errX := Error{Message: msg}

	if err != nil {
		errX.Err = err.Error()
	}

	return errX
}
