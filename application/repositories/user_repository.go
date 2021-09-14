package repositories

import (
	"victorLessa/server/domain"

	"github.com/jinzhu/gorm"
)


type UserRepository interface {
	Insert()
	Index()
	Update()
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo *UserRepositoryDb) Insert(paylod *domain.User) (*domain.User, error) {
	err := repo.Db.Create(paylod).Error

	if err != nil {
		return nil, err
	}
	return paylod, nil
}

func (repo *UserRepositoryDb) Index() (*gorm.DB) {
	var users []domain.User
	
	result := repo.Db.Find(&users)
	
	return result
}

func (repo *UserRepositoryDb) Show(id string) (*gorm.DB) {
	var users []domain.User
	
	var result = repo.Db.Find(&users, "id = ?", id)

	
	return result
}

func(repo *UserRepositoryDb) Update(id string, user *domain.User) (*domain.User, error)  {
		var users []domain.User
		res:= repo.Db.First(&users, "id = ?", id)
		
		err := res.Save(user).Error
		
		if err != nil {
			return nil, err
		}
		
		return user, nil
	}
	
func (repo *UserRepositoryDb) Delete(id string) (*gorm.DB, error) {

	res:= repo.Db.First(&domain.User{}, "id = ?", id)
	
	err := repo.Db.Delete(&domain.User{}, "id = ? ", id).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *UserRepositoryDb) FindByName(username string) *domain.User {
	var users = &domain.User{}
	
	repo.Db.Find(&users, "username = ?", username).Scan(&users)

	return users
}