package offer_usecase

import "github.com/pkg/errors"

var (
	NilDataArg     = errors.New("nil argument in function call")
	MentorNotExist = errors.New("mentor with this ID not exits")
	NotMentor      = errors.New("this user not mentor")
	AlreadyExists  = errors.New("offer from this mentee to this mentor already exists")
	LogicError     = errors.New("mentor = mentee - logic error")
	InvalidOfferID = errors.New("this user have not offer with this id")
)
