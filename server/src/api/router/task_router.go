package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/repo"
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

	apiRouter.GET("/tasks", r.GetTasks)
	apiRouter.GET("/tasks/all", r.GetAllTasks)
	apiRouter.POST("/task", r.CreateTask)
	apiRouter.PUT("/task/:id", r.UpdateTask)
	apiRouter.DELETE("/task/:id", r.DeleteTask)
	apiRouter.POST("/task/revert/:id", r.RevertDeleteTask)
}

func (r *TaskRouter) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* frequency must be at least everyday (1) */
	if task.Frequency < 1 {
		task.Frequency = 1
	}

	result := r.db.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (r *TaskRouter) GetTasks(c *gin.Context) {
	var tasks []model.Task
	r.taskRepository.GetTasks(&tasks)

	c.JSON(http.StatusOK, tasks)
}

func (r *TaskRouter) GetAllTasks(c *gin.Context) {
	var tasks []model.Task
	r.taskRepository.GetAllTasks(&tasks)

	c.JSON(http.StatusOK, tasks)
}

func (r *TaskRouter) UpdateTask(c *gin.Context) {
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

	err = r.taskRepository.UpdateTask(&task, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (r *TaskRouter) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r.taskRepository.DeleteTask(&id)

	c.JSON(http.StatusOK, "Deleted successfully")
}

func (r *TaskRouter) RevertDeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = r.taskRepository.RevertDeleteTask(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "Reverted successfully")
}
