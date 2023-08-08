package utils

type Error struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}
