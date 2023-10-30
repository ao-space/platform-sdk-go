package utils

type Error struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}

func FromError(err error) *Error {
	Err, ok := err.(*Error)
	if !ok {
		return nil
	}
	return Err
}
