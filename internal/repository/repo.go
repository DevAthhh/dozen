package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r *Repository) NewTaskRepo() *TaskRepository {
	return NewTaskRepository(r.db)
}

func (r *Repository) NewUserRepo() *UserRepository {
	return NewUserRepository(r.db)
}

func (r *Repository) NewGroupRepo() *GroupRepository {
	return NewGroupRepository(r.db)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
