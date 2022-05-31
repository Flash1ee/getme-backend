package dto

import (
	"getme-backend/internal/app/offer/entities"
)

type OfferUserUsecaseDTO struct {
}
type OfferUsecaseDTO struct {
	ID        int64
	SkillName string
	MentorID  int64
	MenteeID  int64
}
type OfferAcceptUsecaseDTO struct {
	OfferID     int64
	Title       string
	Description string
	Skills      []string
}

func ToOfferAcceptUsecaseDTO(req RequestAcceptOffer) *OfferAcceptUsecaseDTO {
	return &OfferAcceptUsecaseDTO{
		Title:       req.Title,
		Description: req.Description,
		Skills:      req.Skills,
	}
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
