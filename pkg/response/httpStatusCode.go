package response

const (
	SuccessCode         = 20001
	ErrParamInvalidCode = 20003
	ErrInvalidToken     = 20004
    // Register response codes
    ErrUserAlreadyExists = 50001
)

var msg = map[int]string{
	SuccessCode:         "Success",
	ErrParamInvalidCode: "Invalid email address",
	ErrInvalidToken:     "Invalid token",

    ErrUserAlreadyExists: "User already exists",
}
