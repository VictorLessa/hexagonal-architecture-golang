package usecases

import (
	"victorLessa/server/application/repositories"
	"victorLessa/server/domain"

	"github.com/jinzhu/gorm"
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

func(u *UserUseCase) Index() (*gorm.DB) {
	
	users := u.UserRepository.Index()


	return users
}

func (u *UserUseCase) Show(id string) (*gorm.DB) {
	res := u.UserRepository.Show(id)

	return res
}

func (u *UserUseCase) Update(id string, paylod *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Update(id, paylod)

	if err != nil {

		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) Delete(id string) (*gorm.DB, error) {
	res, err := u.UserRepository.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}