package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/repo"
	"routiner/server/src/util"
)

type TaskRouter struct {
	db             *gorm.DB
	taskRepository *repo.TaskRepository
}

func NewTaskRouter(
	db *gorm.DB,
) *TaskRouter {
	return &TaskRouter{
		db:             db,
		taskRepository: repo.NewTaskRepository(db),
	}
}

func (r *TaskRouter) InitTaskEndpoint(apiRouter *gin.RouterGroup) {

	apiRouter.GET("/task/date", r.GetTasksInDate)
	apiRouter.GET("/task/:id", r.GetTaskByID)
	apiRouter.PUT("/task/:id", r.UpdateTaskByID)
	apiRouter.DELETE("/task/:id", r.DeleteTaskByID)
}

func (r *TaskRouter) GetTaskByID(c *gin.Context) {
	var task model.Task
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r.taskRepository.GetTaskByID(&task, &id)

	c.JSON(http.StatusOK, task)
}

func (r *TaskRouter) GetTasksInDate(c *gin.Context) {
	var tasks []model.Task

	day, dErr := strconv.Atoi(c.Query("d"))
	month, mErr := strconv.Atoi(c.Query("m"))
	year, yErr := strconv.Atoi(c.Query("y"))

	date := util.GetTodayBegin()

	/* if given all 3 parameters, set the date to the given one */
	if dErr == nil && mErr == nil && yErr == nil {
		date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	}

	r.taskRepository.GetTasksInDate(&tasks, &date, false)

	c.JSON(http.StatusOK, tasks)
}

func (r *TaskRouter) UpdateTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var task model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.taskRepository.UpdateTaskByID(&id, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (r *TaskRouter) DeleteTaskByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r.taskRepository.DeleteTaskByID(&id)

	c.JSON(http.StatusOK, "Delete successfully")
}
