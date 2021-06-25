package repositories

import (
	"victorLessa/server/domain"

	"github.com/jinzhu/gorm"
)


type UserRepository interface {
	Insert()
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo *UserRepositoryDb) Insert(payload *domain.User) (*domain.User, error) {
	err := repo.Db.Create(payload).Error

	if err != nil {
		return nil, err
	}
	return payload, nil
}