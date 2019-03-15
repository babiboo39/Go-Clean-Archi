package repository

import (
	"MPPLProject/authentication"
	"MPPLProject/authentication/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct{
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) authentication.RepositoryUser {
	return &UserRepository{Conn}
}

func (ur *UserRepository) Fetch() (res []*models.User, err error) {
	var users []*models.User
	err = ur.Conn.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetById(id uint) (*models.User, error) {
	var user_ models.User
	err := ur.Conn.Find(&user_, id).Error

	if err != nil {
		return nil,err
	}

	return &user_, nil
}

func (ur *UserRepository) Update(u *models.User) error {
	var user_ models.User
	ur.Conn.Find(user_, )

	err := ur.Conn.Save(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Store(u *models.User) error {
	err := ur.Conn.Create(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id uint) error {
	var user_ models.User
	ur.Conn.Find(&user_)
	err := ur.Conn.Delete(&user_).Error

	if err != nil {
		return err
	}

	return nil
}