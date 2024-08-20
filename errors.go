package openapi

type ErrorCode string

var (
	BadRequestErrCode        ErrorCode = "BAD_REQUEST"
	ValidationErrCode        ErrorCode = "VALIDATION"
	UnknownErrCode           ErrorCode = "UNKNOWN"
	MethodNotFoundErrCode    ErrorCode = "METHOD_NOT_FOUND"
	InvalidHttpMethodErrCode ErrorCode = "INVALID_HTTP_METHOD"
)
