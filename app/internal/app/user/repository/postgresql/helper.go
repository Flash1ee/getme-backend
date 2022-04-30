package repository_postgresql

import (
	"testing"
	"time"

	"getme-backend/internal/app/user/entities"
)

func getRepositoryData(t *testing.T) *entities.User {
	t.Helper()

	return &entities.User{
		ID:           1,
		TelegramID:   123456,
		FirstName:    "Vasiliy",
		LastName:     "Alexeev",
		Nickname:     "vasax",
		About:        "this is some information",
		Avatar:       "/img/1.png",
		Email:        "vasyugan@gmai.com",
		IsSearchable: false,
		CreatedAt:    time.Now().Add(-3600),
		UpdatedAt:    time.Now(),
	}

}
