package models

type Account struct {
	Series		uint 		`gorm:"PRIMARY_KEY:AUTO_INCREMENT" json:"series"`

	Token   	string		`json:"token"`
	UserName	string		`gorm:"unique;not null" json:"username"`
	Password	string		`json:"password"`

	UserID		uint32 		`json:"user_id"`
	Users		User 		`json:"user"`
}