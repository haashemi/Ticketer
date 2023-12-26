package api

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewError(msg string, err error) Error {
	errX := Error{Message: msg}

	if err != nil {
		errX.Error = err.Error()
	}

	return errX
}
