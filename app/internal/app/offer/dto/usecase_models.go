package dto

import "getme-backend/internal/app/offer/entities"

type OfferUserUsecaseDTO struct {
}
type OfferUsecaseDTO struct {
	ID        int64
	SkillName string
	MentorID  int64
	MenteeID  int64
}

func ToOfferUsecaseDTO(req RequestCreateOffer) *OfferUsecaseDTO {
	return &OfferUsecaseDTO{
		SkillName: req.SkillName,
		MentorID:  req.MentorID,
	}
}

func (d *OfferUsecaseDTO) ToEntity() *entities.Offer {
	return &entities.Offer{
		SkillName: d.SkillName,
		MentorID:  d.MentorID,
		MenteeID:  d.MenteeID,
	}
}
