package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/repo"
)

type LogRouter struct {
	db            *gorm.DB
	logRepository *repo.LogRepository
}

func NewLogRouter(
	db *gorm.DB,
) *LogRouter {
	return &LogRouter{
		db:            db,
		logRepository: repo.NewLogRepository(db),
	}
}

func (r *LogRouter) InitLogEndpoint(apiRouter *gin.RouterGroup) {

	apiRouter.POST("/log", r.CreateLog)
	apiRouter.GET("/log/:id", r.GetLogByID)
	apiRouter.DELETE("/log/:id", r.DeleteLogByID)
}

func (r *LogRouter) CreateLog(c *gin.Context) {
	var log model.Log
	if err := c.BindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.logRepository.CreateLog(&log)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, log)
}

func (r *LogRouter) GetLogByID(c *gin.Context) {
	var log model.Log
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r.logRepository.GetLogByID(&log, &id)

	c.JSON(http.StatusOK, log)
}

func (r *LogRouter) DeleteLogByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r.logRepository.DeleteLogByID(&id)

	c.JSON(http.StatusOK, "Delete successfully")
}
