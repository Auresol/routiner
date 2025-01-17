package repo

import (
	"errors"

	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/util"
)

type RoutineRepository struct {
	db             *gorm.DB
	taskRepository *TaskRepository
}

func NewRoutineRepository(
	db *gorm.DB,
) *RoutineRepository {
	return &RoutineRepository{
		db:             db,
		taskRepository: NewTaskRepository(db),
	}
}

func (r *RoutineRepository) CreateRoutine(routine *model.Routine) error {

	/* make sure that ActiveDate will always be the beginning of the day */
	routine.ActiveDate = util.GetDateBegin(&routine.ActiveDate)

	result := r.db.Create(&routine)
	if result.Error != nil {
		return result.Error
	}

	//util.Log("Call CreateRoutine")
	r.taskRepository.GenerateTasksForward(routine)

	return nil
}

func (r *RoutineRepository) GetRoutines(routines *[]model.Routine) {
	r.db.Find(&routines)
}

func (r *RoutineRepository) GetAllRoutines(routines *[]model.Routine) {
	r.db.Unscoped().Find(&routines)

}

func (r *RoutineRepository) UpdateRoutine(routine_in *model.Routine, id *int) error {

	// var routine model.Routine
	// result := r.db.First(&routine, id)

	/* TODO:
	if frequency, dayInWeekly, dueIn, activeDate, or forceReset is change (result in occurred date change) do the following ->
	- delete task with status = false and recreate task in the date which is not yet created
	*/

	result := r.db.Model(&model.Routine{}).Where("id = ?", id).Updates(routine_in)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("routine not found")
	}

	return nil
}

func (r *RoutineRepository) DeleteRoutine(id *int) error {

	routine := model.Routine{ID: uint(*id)}
	result := r.db.Find(&routine)
	if result.RowsAffected == 0 {
		return errors.New("no routine found")
	}

	result = r.db.Delete(&routine)
	if result.Error != nil {
		return result.Error
	}

	initialDate := util.GetTodayBegin()

	/* if ActiveDate come after today, start at that date */
	if initialDate.Compare(routine.ActiveDate) == -1 {
		initialDate = routine.ActiveDate
	}

	//util.Debuggers.Log("Info", "initial date for delete: "+initialDate.String())

	result = r.db.Where("routine_id = ? AND begin >= ?", id, initialDate).Delete(&model.Task{})
	//util.Debuggers.Log("Info", "Deleted row affected: "+strconv.Itoa(int(result.RowsAffected)))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *RoutineRepository) RevertDeleteRoutine(id *int) error {

	routine := model.Routine{ID: uint(*id)}
	result := r.db.Find(&routine)
	if result.RowsAffected == 0 {
		return errors.New("no routine found")
	}

	result = r.db.Unscoped().Model(&model.Routine{}).Where("id = ?", id).UpdateColumn("deleted_at", nil)

	if result.Error != nil {
		return result.Error
	}

	initialDate := util.GetTodayBegin()
	/* if ActiveDate come after today, start at that date */
	if initialDate.Compare(routine.ActiveDate) == -1 {
		initialDate = routine.ActiveDate
	}
	result = r.db.Unscoped().Model(&model.Task{}).Where("routine_id = ? AND begin >= ?", id, initialDate).UpdateColumn("deleted_at", nil)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
