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
	taskRepository     *repo.TaskRepository
	logRepository      *repo.LogRepository
	calenderRepository *repo.CalenderRepository
}

func NewMockRouter(
	db *gorm.DB,
) *MockRouter {
	return &MockRouter{
		db:                 db,
		taskRepository:     repo.NewTaskRepository(db),
		logRepository:      repo.NewLogRepository(db),
		calenderRepository: repo.NewCalenderRepository(db),
	}
}

func (r *MockRouter) InitMockEndpoint(mockRouter *gin.RouterGroup) {

	mockRouter.POST("/task", r.MockTask)
	mockRouter.POST("/month", r.MockMonth)
	mockRouter.POST("/clear", r.MockClear)
}

type mockTaskBody struct {
	RoutineOneAmount int `json:"routine_one_amuont"`
	RoutineTwoAmount int `json:"routine_two_amuont"`
}

func (r *MockRouter) MockTask(c *gin.Context) {
	var mockTaskBody mockTaskBody
	if err := c.BindJSON(&mockTaskBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for range mockTaskBody.RoutineOneAmount {
		var task model.Task
		task.Title = gofakeit.Noun()
		task.Explain = gofakeit.SentenceSimple()
		task.IconCodePoint = 57583
		task.RoutineMode = 1
		task.DayInWeekly = int8(gofakeit.Number(1, 63))

		task.Frequency = 1

		task.CreatedAt = time.Now().AddDate(-2, 0, 0)

		r.taskRepository.CreateTask(&task)

	}

	for range mockTaskBody.RoutineTwoAmount {
		var task model.Task
		task.Title = gofakeit.Noun()
		task.Explain = gofakeit.SentenceSimple()
		task.IconCodePoint = 57583
		task.RoutineMode = 2
		task.Frequency = gofakeit.Number(1, 31)
		task.ResetOnMonth = gofakeit.Bool()

		task.CreatedAt = time.Now().AddDate(-1, 0, 0)

		r.taskRepository.CreateTask(&task)
	}

	var tasks []model.Task
	r.taskRepository.GetTasks(&tasks)

	c.JSON(http.StatusOK, tasks)
}

type mockMonthBody struct {
	StartYear  int `json:"start_year"`
	StartMonth int `json:"start_month"`
	Amount     int `json:"amount"`
}

func (r *MockRouter) MockMonth(c *gin.Context) {
	//var tasks []model.Task

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

	var calender model.Calender
	var log model.Log
	finish := 0

	for startDate != endDate {
		var tasks []model.Task
		r.calenderRepository.GetTasksInDate(&tasks, &startDate)

		calender.Date = startDate
		finish = 0

		for _, task := range tasks {

			if gofakeit.Number(1, 10) > 3 {
				finish++
				log.Date = startDate
				log.TaskID = task.ID
				log.Detail = gofakeit.SentenceSimple()
				result := r.db.Create(&log)
				if result.Error != nil {
					task.Log = append(task.Log, log)
				}

				r.db.Save(&task)
			}
		}

		model.CalenderStatusUpdate(&calender, finish, len(tasks))
		r.db.Create(&calender)

		startDate = startDate.AddDate(0, 0, 1)
	}

	c.JSON(http.StatusOK, "Mock finished")
}

func (r *MockRouter) MockClear(c *gin.Context) {

	r.db.Exec("DELETE FROM tasks;")
	r.db.Exec("DELETE FROM logs;")
	r.db.Exec("DELETE FROM calenders;")

	c.JSON(http.StatusOK, "Mock cleared")
}
