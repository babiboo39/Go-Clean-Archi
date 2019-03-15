package models

type Role struct {
	ID uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id_role"`
	Role string `gorm:"varchar(100)" json:"role"`

}
