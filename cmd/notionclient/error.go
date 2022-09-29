package notionclient

type ErrorCode string

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type Error struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
