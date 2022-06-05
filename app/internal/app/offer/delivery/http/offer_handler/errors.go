package offer_handler

import "github.com/pkg/errors"

var (
	MentorNotFound     = errors.New("mentor not found")
	SkillNotFound      = errors.New("skill with this name not found")
	OfferAlreadyExists = errors.New("offer from this mentee to this mentor alredy exists")
	UserNotMentor      = errors.New("this user not mentor")
	LogicError         = errors.New("logic error")
)
