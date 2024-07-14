package testdata

import (
	"inno/attestation1/internal/entity"
	"os"
)

func GenerateUserRequest(userId, token string, file *os.File) entity.Message {
	message := entity.Message{
		Token:  token,
		FileID: file.Name(),
		Data:   userId,
	}
	return message
}
