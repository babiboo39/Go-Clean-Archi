package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	FirstName			string `json:"first_name"`
	LastName			string `json:"last_name"`
	Email       		string `gorm:"type:varchar(100);unique;" json:"email"`
	Password			string `json:"password"`
	PhoneNumber			string `json:"phone_number"`
	ActivationStatus	uint64 `json:"activation_status"`
}