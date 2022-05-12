package usecase

import (
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
)

type BaseUsecase struct {
}

func (u *BaseUsecase) GenToken(initValue string) string {
	value := append([]byte(initValue), uuid.NewV4().Bytes()...)
	hash := sha3.Sum512(value)

	return hex.EncodeToString(hash[:])
}
