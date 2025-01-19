package biz_error

type BizError struct {
	code    int64
	message string
}

func (e *BizError) Error() string {
	return e.message
}

var (
	SignUpError = newBizError(1000, "sign up failed, email has been used")
)

func newBizError(code int64, message string) *BizError {
	return &BizError{
		code:    code,
		message: message,
	}
}
