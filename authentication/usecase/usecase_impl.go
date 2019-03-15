package usecase

import (
	"MPPLProject/authentication"
	"MPPLProject/authentication/models"
)

type userUseCase struct {
	userRepo authentication.RepositoryUser
}

func(uuc *userUseCase) Fetch() (res []*models.User, err error) {
	res, err = uuc.userRepo.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func(uuc *userUseCase) GetById(id uint) (*models.User, error) {
	res, err := uuc.userRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uuc *userUseCase) Update(u *models.User) error {
	return uuc.userRepo.Update(u)
}

func (uuc *userUseCase) Store(u *models.User) error {
	return uuc.userRepo.Store(u)
}

func (uuc *userUseCase) Delete(id uint) error {
	return uuc.userRepo.Delete(id)
}

func NewUserUseCase(repository authentication.RepositoryUser) authentication.UseCase {
	return &userUseCase{userRepo:repository}
}