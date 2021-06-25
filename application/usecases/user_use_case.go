package usecases

import (
	"victorLessa/server/application/repositories"
	"victorLessa/server/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepositoryDb
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Insert(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}