package models

type UserStatus struct {
	IdStatusUser 	uint	`gorm:"Primary_Key;Auto_Increment" json:"id_status_user"`
	Series			uint	`json:"series"`
}
