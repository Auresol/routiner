package repo

import (
	"time"

	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/util"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(
	db *gorm.DB,
) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (r *LogRepository) CreateLog(log *model.Log) error {

	log.Date = util.GetDateBegin(&log.Date)

	result := r.db.Create(&log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *LogRepository) GetLogByID(log *model.Log, id *int) {

	r.db.Model(log).Where("id = ?", id)
}

func (r *LogRepository) DeleteLogByID(id *int) error {

	result := r.db.Delete(&model.Log{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/* Callable from within only */
func (r *LogRepository) DeleteLog(Task *model.Task, date time.Time) error {

	date = util.GetDateBegin(&date)
	result := r.db.Where("Task_id = ? AND Date = ?", Task.ID, date).Delete(&model.Log{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
