package models

import (
	"App-Name/config"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type SQLModel struct {
	// Good for Postgres
	// for MySQL, remove first 3 fields and use gorm.Model instead
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DBResponse `gorm:"-" json:"-"`
}

type DBResponse struct {
	Message string
	Status  int
}

func (s *SQLModel) HandleErr(result *gorm.DB, logger *log.Logger) {
	s.DBResponse.Message, s.DBResponse.Status = s.HandleDbError(result, logger)
	if s.Status != 0 {
		panic("Database Error")
	}
}

func (s *SQLModel) HandleDbError(result *gorm.DB, logger *log.Logger) (string, int) {
	//Handle common errors
	if result.Error != nil {
		logger.Println("Database Error: ", result.Error)
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return "Records not found", 404
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return "Record Exists", 400
		default:
			//Catch all other errors
			return "Internal Server Error, Kindly Contact Support", 500
		}
	}
	return "", 0
}

func (s *SQLModel) StoreBB(modelInstance, payload interface{}, logger *log.Logger) {
	result := config.DBClient.Model(modelInstance).Create(payload)
	s.HandleErr(result, logger)
}
