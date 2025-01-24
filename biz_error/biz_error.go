package biz_error

type BizError struct {
	code    int64
	message string
}

func (e *BizError) Error() string {
	return e.message
}

var (
	SignUpError          = newBizError(1000, "sign up failed, email has been used")
	UpdatePasswordError  = newBizError(1002, "update password failed, old password is same as new password")
	UpdatePasswordError2 = newBizError(1003, "update password failed, old password is wrong")
	LoginError           = newBizError(1004, "login failed, email is not exist")
	LoginError2          = newBizError(1005, "login failed, password is wrong")
	TokenError           = newBizError(1006, "token convert to claims failed")
)

func newBizError(code int64, message string) *BizError {
	return &BizError{
		code:    code,
		message: message,
	}
}
