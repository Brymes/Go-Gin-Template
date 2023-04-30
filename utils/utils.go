package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"runtime/debug"
	"strings"
)

func GenerateUniqueID(length int, charset string) string {
	bytes := make([]byte, length)

	var chars string

	switch charset {
	case "alphabets":
		chars = "abcdefghijklmnopqrstuvwxyz"
	case "numeric":
		chars = "0123456789"
	case "alphanumeric":
		chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	default:
		chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	}

	if _, err := rand.Read(bytes); err != nil {
		panic("Internal Server Error whilst Generating password")
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes)
}

func OnErrorPanic(err error, helpText string) {
	if err != nil {
		panic(fmt.Sprintf("%s: \n, %v", helpText, err))
	}
}

func OnErrorLog(err error) {

}

func HandleDbError(result *gorm.DB, logger *log.Logger) (string, int) {
	if result.Error != nil {
		logger.Println("Database Error: ", result.Error)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//Handle Records Not Found

			return "Internal Server Error: Record(s) not Found", 404
		} else if strings.Contains(result.Error.Error(), "Error 1062") {
			//Handle Duplicate Key Error for Mariadb

			fieldName := strings.Split(result.Error.Error(), "for key")
			return fmt.Sprintf("Internal Server Error: Field %s already exists", fieldName[1]), 400
		} else {
			//Catch all other errors

			return "Internal Server Error, Kindly Contact Support", 500
		}
	}

	return "", 0
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		panic("Internal Server Error whilst Hashing password")
	}

	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HandlePanicMacro(err interface{}, logger *log.Logger) bool {
	if err != nil {
		if logger == nil {
			logger = log.Default()
		}
		logger.Println(err)
		logger.Println(string(debug.Stack()))
		return true
	}
	return false
}
