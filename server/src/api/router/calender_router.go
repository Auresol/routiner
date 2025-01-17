package router

import (
	"net/http"
	"routiner/server/src/model"
	"routiner/server/src/repo"
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

	apiRouter.GET("/cal/month", r.GetMonthSummary)
}

func (r *CalenderRouter) GetMonthSummary(c *gin.Context) {
	var calender []model.Calender

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
