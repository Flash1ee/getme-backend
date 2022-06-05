package auth_usecase

import "github.com/pkg/errors"

var (
	EmptyPassword            = errors.New("empty password")
	ArgError                 = errors.New("invalid argument, expected not nil")
	BadAuth                  = errors.New("authorization error - not valid data")
	IncorrectLoginOrPassword = errors.New("incorrect login or password")
	SimpleAuthExists         = errors.New("simple auth exists")
	UserExist                = errors.New("user already exist")
	InvalidOldNickname       = errors.New("old nickname not equal user nickname with this users_id")
	NicknameExists           = errors.New("this nickname already exist")
	LoginExists              = errors.New("this login already exist")

	BadEncrypt           = errors.New("unsuccessful encrypt user")
	OldPasswordEqualNew  = errors.New("the new password must be different from the old one")
	IncorrectNewPassword = errors.New("new password not valid")
)
