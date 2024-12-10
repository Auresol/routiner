package router

import (
	"net/http"
	"routiner/server/src/model"
	"routiner/server/src/repo"
	"routiner/server/src/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CalenderRouter struct {
	db              *gorm.DB
	calenderHandler *repo.CalenderRepository
}

func NewCalenderRouter(
	db *gorm.DB,
) *CalenderRouter {
	return &CalenderRouter{
		db:              db,
		calenderHandler: repo.NewCalenderRepository(db),
	}
}
func (r *CalenderRouter) InitCalenderEndpoint(apiRouter *gin.RouterGroup) {

	apiRouter.GET("/cal/date", r.GetTasksInDate)
	apiRouter.GET("/cal/month", r.GetMonthSummary)
	apiRouter.GET("/cal/date/sum", r.GetDaySummary)
}

func (r *CalenderRouter) GetTasksInDate(c *gin.Context) {
	var tasks []model.Task

	day, dErr := strconv.Atoi(c.Query("d"))
	month, mErr := strconv.Atoi(c.Query("m"))
	year, yErr := strconv.Atoi(c.Query("y"))

	date := util.GetTodayBegin()

	/* if given all 3 parameters, set the date to the given one */
	if dErr == nil && mErr == nil && yErr == nil {
		date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	}

	r.calenderHandler.GetTasksInDate(&tasks, &date)

	c.JSON(http.StatusOK, tasks)
}

func (r *CalenderRouter) GetMonthSummary(c *gin.Context) {
	var calender model.Calender

	month, mErr := strconv.Atoi(c.Query("m"))
	year, yErr := strconv.Atoi(c.Query("y"))

	if mErr != nil || yErr != nil {
		year = time.Now().Year()
		month = int(time.Now().Month())
	}

	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)

	r.calenderHandler.GetMonthSummary(&calender, &date)

	c.JSON(http.StatusOK, calender)
}

func (r *CalenderRouter) GetDaySummary(c *gin.Context) {
	var tasks []model.Task
	var logs []model.Log
	var calender model.Calender

	day, dErr := strconv.Atoi(c.Query("d"))
	month, mErr := strconv.Atoi(c.Query("m"))
	year, yErr := strconv.Atoi(c.Query("y"))

	date := util.GetTodayBegin()

	/* if given all 3 parameters, set the date to the given one */
	if dErr == nil && mErr == nil && yErr == nil {
		date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	}

	r.calenderHandler.GetDaySummary(&tasks, &logs, &calender, &date)

	c.JSON(http.StatusOK, gin.H{"calender": calender, "logs": logs, "tasks": tasks})
}
