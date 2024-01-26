package dto

import (
	"jade_backend/utils/stacktrace"
)

const (
	NoErr           stacktrace.ErrorCode = iota
	ErrEmailExist   stacktrace.ErrorCode = 1
	ErrCaptcha      stacktrace.ErrorCode = 2
	ErrPassword     stacktrace.ErrorCode = 3
	ErrUserNotFound stacktrace.ErrorCode = 4

	ErrShortLinkExist stacktrace.ErrorCode = 5
	ErrNoJade         stacktrace.ErrorCode = 6
	ErrPrivilege      stacktrace.ErrorCode = 7

	ErrShortLinkActive stacktrace.ErrorCode = 8
	ErrShortLinkTime   stacktrace.ErrorCode = 9

	BadRequest    stacktrace.ErrorCode = 400
	InternalError stacktrace.ErrorCode = 500
)
