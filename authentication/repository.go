package authentication

import models "MPPLProject/authentication/models"

type RepositoryUser interface {
	Fetch() (res []*models.User, err error)
	GetById(id uint) (*models.User, error)
	Update(u *models.User) error
	Store(u *models.User) error
	Delete(id uint) error
}