package models

import (
  "time"
  "gorm.io/gorm"
)

type BaseModel struct {
  ID        uint       `gorm:"primary_key" json:"id"`
  CreatedAt time.Time  `json:"created_at"`
  UpdatedAt time.Time  `json:"updated_at"`
}

type Todo struct {
  BaseModel
	Title  string  `json:"title"`
	Author string  `json:"author"`
}

type TodoModel struct {
  DB *gorm.DB
}

func NewTodoModel(db *gorm.DB) *TodoModel {
  return &TodoModel{DB: db}
}

func (m *TodoModel) GetAll() ([]Todo, error) {
  var todos []Todo
  if err := m.DB.Find(&todos).Error; err != nil {
    return nil, err
  }

  return todos, nil
}

func (m *TodoModel) GetByID(id uint) (Todo, error) {
  var todo Todo
  if err := m.DB.Where("id = ?", id).First(&todo).Error; err != nil {
    return Todo{}, err
  }

  return todo, nil
}

func (m *TodoModel) Create(title string, author string) (Todo, error) {
	todo := Todo{Title: title, Author: author}
	if err := m.DB.Create(&todo).Error; err != nil {
				return Todo{}, err
	}

	return todo, nil
}

func (m *TodoModel) Update(id uint, title string, author string) (Todo, error) {
	todo, err := m.GetByID(id)
	if err != nil {
		return Todo{}, err
	}

	if err := m.DB.Model(&todo).Updates(map[string]interface{}{"title": title, "author": author}).Error; err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (m *TodoModel) Delete(id uint) error {
	todo, err := m.GetByID(id)
	if err != nil {
				return err
	}

	return m.DB.Delete(&todo).Error
}
