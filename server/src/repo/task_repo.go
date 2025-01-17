package repo

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/util"
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

	result := r.db.Create(task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepository) GetTaskByID(task *model.Task, id *int) {

	r.db.Model(task).Where("id = ?", id)
}

/* Get all tasks in the given date, will generate new task if the given date is not being generated */
func (r *TaskRepository) GetTasksInDate(tasks *[]model.Task, date *time.Time, forceRegenerate bool) {

	var calender model.Calender
	result := r.db.Find(&calender, "date = ?", *date)

	if result.RowsAffected == 0 || !calender.TaskIsCreated || forceRegenerate {
		calender.Date = *date
		r.RecreateTasksInDate(date)
		calender.TaskIsCreated = true
	}

	r.CountAndUpdateCalender(&calender)
	r.db.Save(&calender)

	r.db.Where("? BETWEEN begin AND due", *date).Preload("Routine").Find(&tasks)

}

func (r *TaskRepository) UpdateTaskByID(id *int, task_in *model.Task) error {

	var task model.Task
	result := r.db.First(&task, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	if task.Status != task_in.Status {
		r.SingleTaskUpdate(&task)

	}

	r.db.Model(&model.Task{}).Where("id = ?", id).Updates(task_in)
	return nil

}

func (r *TaskRepository) DeleteTaskByID(id *int) error {

	result := r.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
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

/* Generate task for certain date */
func (r *TaskRepository) RecreateTasksInDate(date *time.Time) {

	var routines []model.Routine

	/* remove old task begin with this date */
	r.db.Unscoped().Where("begin = ?", *date).Delete(&model.Task{})

	r.db.Find(&routines)

	for _, routine := range routines {

		r.GenerateTaskFromRoutineInDate(&routine, date)
	}

}

/* Calculate the number of finished task compared to total task */
func (r *TaskRepository) CountAndUpdateCalender(calender *model.Calender) {

	var tasks []model.Task
	r.db.Where("begin = ?", calender.Date).Preload("Routine").Find(&tasks)
	calender.TotalTasks = len(tasks)
	finished := 0
	for _, element := range tasks {
		if element.Status {
			finished++
		}
	}
	calender.FinishedTasks = finished

	r.updateCalenderStatus(calender)
}

/* Update finished percentage of calender based on task begin and due */
func (r *TaskRepository) SingleTaskUpdate(task *model.Task) {
	var calenders []model.Calender

	r.db.Model(&model.Calender{}).Where("date BETWEEN ? AND ?", task.Begin, task.Due).Find(&calenders)
	for _, element := range calenders {
		if task.Status {
			element.FinishedTasks++
		} else {
			element.FinishedTasks--
		}
		r.updateCalenderStatus(&element)
	}
}

/* Update calender status based on its total and finished */
func (r *TaskRepository) updateCalenderStatus(calender *model.Calender) {

	if calender.TotalTasks == 0 {
		calender.Status = model.NONE
		return
	}

	percent := float32(calender.FinishedTasks) / float32(calender.TotalTasks)

	if percent > 1 {
		calender.Status = model.EXTRA
	} else if percent == 1 {
		calender.Status = model.COMPLETE
	} else if percent >= 0.7 {
		calender.Status = model.NEARLY
	} else if percent > 0 {
		calender.Status = model.SOME
	} else {
		calender.Status = model.NONE
	}
}

/* Create new task for date that already have calender */
func (r *TaskRepository) GenerateTasksForward(routine *model.Routine) {
	initialDate := util.GetTodayBegin()

	/* if ActiveDate come after today, start at that date */
	if initialDate.Compare(routine.ActiveDate) == -1 {
		initialDate = routine.ActiveDate
	}

	var calenders []model.Calender
	r.db.Model(&model.Calender{}).Where("date >= ?", initialDate).Find(&calenders)

	if len(calenders) == 0 {
		return
	}

	lastDate := calenders[len(calenders)-1].Date

	if routine.RoutineMode == model.WEEKLY || routine.RoutineMode == model.PERIOD {

		for initialDate.Compare(lastDate) != 1 {
			r.GenerateTaskFromRoutineInDate(routine, &initialDate)
			initialDate = initialDate.AddDate(0, 0, 1)
		}

	} else if routine.RoutineMode == model.TODO {
		if routine.ActiveDate.Compare(lastDate) != 1 {
			r.GenerateTaskFromRoutineInDate(routine, &routine.ActiveDate)
		}
	}
}

/* Check if this routine is belong to the date, then created if it's belong */
func (r *TaskRepository) GenerateTaskFromRoutineInDate(routine *model.Routine, date *time.Time) {

	dayInWeek := int(date.Weekday())

	var days [7]bool
	var nextDate time.Time
	var isCreated = false

	/* Check if the routine is actived compared to the date */
	if routine.ActiveDate.Compare(*date) == 1 {
		return
	}

	if routine.RoutineMode == model.WEEKLY {

		days = model.BitmaskDecoding(routine.DayInWeekly)
		if days[dayInWeek] {
			isCreated = true

			/* find the next routine */
			var i = 1
			for !days[(dayInWeek+i)%7] {
				i++
			}

			nextDate = date.AddDate(0, 0, i)
		}

	} else if routine.RoutineMode == model.PERIOD {
		nextDate = date.AddDate(0, 0, routine.Frequency)

		if routine.ResetOnMonth && (date.Day()-1)%routine.Frequency == 0 {

			isCreated = true

			/* if the next routine surpass this month, it will be the beginning of next month */
			if nextDate.Month() > date.Month() || nextDate.Year() > date.Year() {
				nextDate = util.GetMonthBegin(date).AddDate(0, 1, 0)
			}

		} else if int(date.Sub(routine.ActiveDate).Hours()/24)%routine.Frequency == 0 {

			isCreated = true
			nextDate = date.AddDate(0, 0, routine.Frequency)
		}

	} else if routine.RoutineMode == model.TODO && routine.ActiveDate.Compare(*date) == 0 {
		isCreated = true
		nextDate = date.AddDate(0, 0, routine.DueIn-1)
	}

	if !isCreated {
		return
	}

	/* due date is the day BEFORE next routine */
	nextDate = nextDate.AddDate(0, 0, -1)

	/* due date is the day BEFORE next routine */
	var dueDate = date.AddDate(0, 0, routine.DueIn-1)

	/* due date will be next routine if this routine is ForceReset */
	if routine.ForceReset && dueDate.After(nextDate) {
		dueDate = nextDate
	}

	task := model.Task{
		RoutineID: routine.ID,
		Begin:     *date,
		Due:       dueDate,
		Status:    false,
	}

	r.CreateTask(&task)
}
