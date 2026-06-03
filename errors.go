package openapi

type ErrorCode string

var (
	InvalidDataErrCode       ErrorCode = "INVALID_DATA"
	ValidationErrCode        ErrorCode = "VALIDATION"
	UnknownErrCode           ErrorCode = "UNKNOWN"
	MethodNotFoundErrCode    ErrorCode = "METHOD_NOT_FOUND"
	InvalidHttpMethodErrCode ErrorCode = "INVALID_HTTP_METHOD"
)
