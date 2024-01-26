package dto

import (
	"jade_backend/utils/stacktrace"
)

const (
	NoErr           stacktrace.ErrorCode = iota
	//ErrEmailExist   stacktrace.ErrorCode = 1
	//ErrCaptcha      stacktrace.ErrorCode = 2
	//ErrPassword     stacktrace.ErrorCode = 3
	//ErrUserNotFound stacktrace.ErrorCode = 4

	ErrNoJade			stacktrace.ErrorCode = 1
	ErrJadeExist 		stacktrace.ErrorCode = 2
	ErrJadeNotFound    	stacktrace.ErrorCode = 3
	ErrPrivilege    	stacktrace.ErrorCode = 4

	//ErrShortLinkActive stacktrace.ErrorCode = 8
	//ErrShortLinkTime   stacktrace.ErrorCode = 9

	BadRequest    stacktrace.ErrorCode = 400
	InternalError stacktrace.ErrorCode = 500
)
