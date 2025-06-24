package response

const (
	SuccessCode         = 20001
	ErrParamInvalidCode = 20003
)

var msg = map[int]string{
	SuccessCode:         "Success",
	ErrParamInvalidCode: "Invalid email address",
}
