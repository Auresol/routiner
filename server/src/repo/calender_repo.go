package repo

import (
	"time"

	"gorm.io/gorm"

	"routiner/server/src/model"
)

type CalenderRepository struct {
	db *gorm.DB
}

func NewCalenderRepository(
	db *gorm.DB,
) *CalenderRepository {
	return &CalenderRepository{
		db: db,
	}
}

func (r *CalenderRepository) GetTasksInDate(tasks *[]model.Task, date *time.Time) {

	isNewMonth := date.Day() == 1
	dayInWeek := date.Day() % 7

	var allTasks []model.Task
	r.db.Find(&allTasks)

	for _, element := range allTasks {
		if element.CreatedAt.Compare(*date) == -1 {
			if (element.RoutineMode == model.WEEKLY && model.BitmaskDecoding(element.DayInWeekly, &dayInWeek)) ||
				(element.RoutineMode == model.PERIOD && (element.ResetOnMonth && isNewMonth) || ((date.Day()-element.CreatedAt.Day())%element.Frequency == 0)) {

				*tasks = append(*tasks, element)
			}
		}
	}

}

func (r *CalenderRepository) GetMonthSummary(calender *model.Calender, date *time.Time) {

	r.db.Model(&model.Calender{}).Where("Date BETWEEN ? AND ?", date, date.AddDate(0, 1, 0)).Find(&calender)

}

func (r *CalenderRepository) GetDaySummary(tasks *[]model.Task, logs *[]model.Log, calender *model.Calender, date *time.Time) {

	r.db.Model(&model.Calender{}).Where("Date = ?", date).Find(&calender)
	r.db.Model(&model.Log{}).Where("Date = ?", date).Find(&logs)
	r.GetTasksInDate(tasks, date)

}
