package framework

type ErrorCode string

const (
	AuthenticationError ErrorCode = "auth_error"
	UnexpectedError     ErrorCode = "unexpected_error"
	UserExistsError     ErrorCode = "user_exists"
	InvalidPayload      ErrorCode = "invalid_payload"
)

type Error struct {
	ErrorCode ErrorCode `json:"error_code"`
	Error     string    `json:"error"`
}

func NewError(errorCode ErrorCode, err string) Error {
	return Error{errorCode, err}
}
