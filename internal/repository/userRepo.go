package repository

import (
	"github.com/DevAthhh/DoZen/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) CreateUser(username, email, pwd string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: string(passwordHash),
		Groups:   []models.Group{{Name: "self"}},
	}

	return u.db.Create(&user).Error
}

func (u *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := u.db.Preload("Groups.Tasks").Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Preload("Groups.Tasks").Find(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
