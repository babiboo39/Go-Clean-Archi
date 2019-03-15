package authentication

import "MPPLProject/authentication/models"

type UseCase interface {
	Fetch() (res []*models.User, err error)
	GetById(id uint) (*models.User, error)
	Update(u *models.User) error
	Store(u *models.User) error
	Delete(id uint) error
}