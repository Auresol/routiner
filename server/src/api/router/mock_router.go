package router

import (
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/repo"
)

type MockRouter struct {
	db                 *gorm.DB
	routineRepository  *repo.RoutineRepository
	taskRepository     *repo.TaskRepository
	calenderRepository *repo.CalenderRepository
}

func NewMockRouter(
	db *gorm.DB,
) *MockRouter {
	var calenderRepository = repo.NewCalenderRepository(db)
	return &MockRouter{
		db:                 db,
		routineRepository:  repo.NewRoutineRepository(db),
		taskRepository:     repo.NewTaskRepository(db),
		calenderRepository: calenderRepository,
	}
}

func (r *MockRouter) InitMockEndpoint(mockRouter *gin.RouterGroup) {

	mockRouter.POST("/routine", r.MockRoutine)
	mockRouter.POST("/month", r.MockMonth)
	mockRouter.POST("/clear", r.MockClear)
}

type mockRoutineBody struct {
	RoutineOneAmount int `json:"routine_one_amuont"`
	RoutineTwoAmount int `json:"routine_two_amuont"`
}

func (r *MockRouter) MockRoutine(c *gin.Context) {
	var mockRoutineBody mockRoutineBody
	if err := c.BindJSON(&mockRoutineBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for range mockRoutineBody.RoutineOneAmount {
		var routine model.Routine
		routine.Title = gofakeit.Noun()
		routine.Explain = gofakeit.SentenceSimple()
		routine.IconCodePoint = 57583

		routine.RoutineMode = model.WEEKLY
		routine.DayInWeekly = int8(gofakeit.Number(1, 63))
		routine.ForceReset = true
		routine.ActiveDate = time.Now().AddDate(-1, 0, 0)
		routine.DueIn = gofakeit.Number(1, 50)

		routine.Frequency = 1
		r.routineRepository.CreateRoutine(&routine)

	}

	for range mockRoutineBody.RoutineTwoAmount {
		var routine model.Routine
		routine.Title = gofakeit.Noun()
		routine.Explain = gofakeit.SentenceSimple()
		routine.IconCodePoint = 57583

		routine.RoutineMode = model.PERIOD
		routine.Frequency = gofakeit.Number(1, 50)
		routine.ResetOnMonth = gofakeit.Bool()
		routine.ForceReset = true

		routine.ActiveDate = time.Now().AddDate(-1, 0, 0)
		routine.DueIn = gofakeit.Number(1, 50)

		r.routineRepository.CreateRoutine(&routine)
	}

	var routines []model.Routine
	r.routineRepository.GetRoutines(&routines)

	c.JSON(http.StatusOK, routines)
}

type mockMonthBody struct {
	StartYear  int `json:"start_year"`
	StartMonth int `json:"start_month"`
	Amount     int `json:"amount"`
}

func (r *MockRouter) MockMonth(c *gin.Context) {
	//var routines []model.Routine

	var mockMonthBody mockMonthBody
	if err := c.BindJSON(&mockMonthBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if mockMonthBody.Amount == 0 {
		mockMonthBody.Amount = 1
	}

	startDate := time.Date(mockMonthBody.StartYear, time.Month(mockMonthBody.StartMonth), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, mockMonthBody.Amount, 0)

	finish := 0

	for startDate != endDate {
		var tasks []model.Task
		r.taskRepository.GetTasksInDate(&tasks, &startDate, false)

		finish = 0

		for _, task := range tasks {

			if gofakeit.Number(1, 10) > 3 && !task.Status {
				finish++
				task.Status = true
				r.db.Save(&task)
				r.taskRepository.SingleTaskUpdate(&task)
			}
		}

		startDate = startDate.AddDate(0, 0, 1)
	}

	c.JSON(http.StatusOK, "Mock finished")
}

func (r *MockRouter) MockClear(c *gin.Context) {

	r.db.Exec("DELETE FROM routines;")
	r.db.Exec("DELETE FROM tasks;")
	r.db.Exec("DELETE FROM calenders;")

	c.JSON(http.StatusOK, "Mock cleared")
}
