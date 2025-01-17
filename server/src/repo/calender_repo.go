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

/*
func (r *CalenderRepository) GetRoutinesInDate(routines *[]model.Routine, date *time.Time) {

	isNewMonth := date.Day() == 1
	dayInWeek := date.Day() % 7

	var allRoutines []model.Routine
	r.db.Find(&allRoutines)

	for _, element := range allRoutines {
		if element.CreatedAt.Compare(*date) == -1 {
			if (element.RoutineMode == model.WEEKLY && model.BitmaskDecoding(element.DayInWeekly, &dayInWeek)) ||
				(element.RoutineMode == model.PERIOD && (element.ResetOnMonth && isNewMonth) || ((date.Day()-element.CreatedAt.Day())%element.Frequency == 0)) {

				*routines = append(*routines, element)
			}
		}
	}
}
*/

/* Get montly summary */
func (r *CalenderRepository) GetMonthSummary(calender *[]model.Calender, date *time.Time) {

	r.db.Model(&model.Calender{}).Where("date BETWEEN ? AND ?", date, date.AddDate(0, 1, 0)).Find(&calender)

}
