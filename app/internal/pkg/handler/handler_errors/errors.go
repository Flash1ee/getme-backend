package handler_errors

import (
	"errors"
)

/// NOT FOUND
var (
	UserNotFound  = errors.New("user not found")
	OfferNotFound = errors.New("offer not found")
	UnknownRole   = errors.New("unknown role, only mentor/mentee supported")

	UserWithNicknameNotFound = errors.New("user with this nickname not found")
)

/// Fields Incorrect
var (
	InvalidNickname          = errors.New("invalid creator nickname")
	EmptyName                = errors.New("empty name in request")
	IncorrectLoginOrPassword = errors.New("incorrect login or password")
	UserAlreadyExists        = errors.New("user already exists")
	IncorrectNewPassword     = errors.New("invalid new password")
)

// BD Error
var (
	NicknameAlreadyExist = errors.New("nickname already exist")
	BDError              = errors.New("can not do bd operation")
)

// Request Error
var (
	InvalidBody          = errors.New("invalid body in request")
	InvalidParameters    = errors.New("invalid parameters")
	InvalidQueries       = errors.New("invalid parameters in query")
	FileSizeError        = errors.New("size of file very big")
	InvalidFormFieldName = errors.New("invalid form field name for load file")
)

// Session Error
var (
	ErrorCreateSession = errors.New("can not create session")
	DeleteCookieFail   = errors.New("can not delete cookie from session store")
)

// jwt errors
var (
	TokenInvalid = errors.New("invalid token")
)

var InternalError = errors.New("server error")
var NoModify = errors.New("content not modify")

// Logic Error
var (
	LogicError = errors.New("logic error in workking of service")
)
