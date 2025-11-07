package dal

import (
	"fiber-rest/database"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool `gorm:"default:false"`
}

func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(&todo)
}

func GetTodos(dest any) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest)
}

func GetTodoByID(todoID any, dest any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", todoID).First(dest)
}

func UpdateTodo(todoID any, data any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", todoID).Updates(data)
}

func DeleteTodo(todoID any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", todoID).Delete(&Todo{})
}
