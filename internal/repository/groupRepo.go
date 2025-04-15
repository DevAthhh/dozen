package repository

import (
	"fmt"

	"github.com/DevAthhh/DoZen/internal/models"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func (g *GroupRepository) CreateGroup(userIDs []int, name string) error {
	var users []models.User
	for _, id := range userIDs {
		var user models.User
		if err := g.db.First(&user, "id = ?", id).Error; err != nil {
			break
		}
		if user.ID == 0 {
			break
		}
		users = append(users, user)
	}
	fmt.Println(users)
	group := models.Group{
		Name:  name,
		Users: users,
	}

	return g.db.Create(&group).Error
}

func (g *GroupRepository) GetAllTasks(id int) []models.Task {
	return []models.Task{}
}

func (g *GroupRepository) GetAllMembers() []models.User {
	return []models.User{}
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}
