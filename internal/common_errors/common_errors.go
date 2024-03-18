package common_errors

type ErrorCode string

const (
	ErrNotFound ErrorCode = "NOT_FOUND"
	ErrInvalidInput ErrorCode = "INVALID_INPUT"
)

type CommonErrors struct {
	Code ErrorCode
	Message string
}

func (e *CommonErrors) Error() string {
	return e.Message
}

func NewCommonErrors(code ErrorCode, message string) *CommonErrors {
	return &CommonErrors{
		Code: code,
		Message: message,
	}
} 