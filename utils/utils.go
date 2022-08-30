package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"runtime/debug"
)

//House Utility Functions in a central file or special files

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"-"`
	Data    interface{} `json:"data"`
}

func GenerateUniqueID(length int) (string, error) {
	bytes := make([]byte, length)

	chars := "0123456789abcdefghijklmnopqrstuvwxyz"

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

func HandlePanic(response *APIResponse, logger *log.Logger) {
	if err := recover(); err != nil {
		logger.Println(string(debug.Stack()))
		response.Success, response.Code = false, 500
		response.Message = fmt.Sprintf("%v", err)
	}
}
