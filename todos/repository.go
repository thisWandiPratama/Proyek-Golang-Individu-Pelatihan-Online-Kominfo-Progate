package todos

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Todo, error)
	Save(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(id int) (Todo, error)
	FindByID(ID int) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Todo, error) {
	var todo []Todo

	err := r.db.Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) Save(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) Update(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) Delete(id int) (Todo, error) {
	var todo Todo
	err := r.db.Delete(&Todo{}, id).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) FindByID(ID int) (Todo, error) {
	var todo Todo

	err := r.db.Where("id = ?", ID).Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}
