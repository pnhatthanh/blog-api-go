package repositories

import (
	"blogapi/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User)
	Update(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(id string) (*models.User, error)
	Delete(id string) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
func (repo *userRepository) Create(user *models.User) {
	repo.db.Create(user)
}

func (repo *userRepository) Update(user *models.User) (*models.User, error) {
	if err := repo.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (*models.User, error) {
	var _user models.User
	if err := repo.db.Find(&_user, "email=?", email).Error; err != nil {
		return nil, err
	}
	return &_user, nil
}

func (repo *userRepository) GetById(id string) (*models.User, error) {
	var _user models.User
	if err := repo.db.Find(&_user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &_user, nil
}

func (repo *userRepository) Delete(id string) error {
	var _user models.User
	if err := repo.db.Find(&_user, "id=?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Delete(&_user).Error; err != nil {
		return err
	}
	return nil
}
