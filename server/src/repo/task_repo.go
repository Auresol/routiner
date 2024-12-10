package repo

import (
	"gorm.io/gorm"

	"errors"
	"routiner/server/src/model"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(
	db *gorm.DB,
) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) CreateTask(task *model.Task) error {

	result := r.db.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepository) GetTasks(tasks *[]model.Task) {
	r.db.Find(&tasks)
}

func (r *TaskRepository) GetAllTasks(tasks *[]model.Task) {
	r.db.Unscoped().Find(&tasks)

}

func (r *TaskRepository) UpdateTask(task *model.Task, id *int) error {

	result := r.db.Model(&model.Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *TaskRepository) DeleteTask(id *int) error {

	result := r.db.Delete(&model.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepository) RevertDeleteTask(id *int) error {

	result := r.db.Unscoped().Model(&model.Task{}).Where("id = ?", id).UpdateColumn("deleted_at", nil)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
