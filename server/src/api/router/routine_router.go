package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"routiner/server/src/model"
	"routiner/server/src/repo"
)

type RoutineRouter struct {
	db                *gorm.DB
	routineRepository *repo.RoutineRepository
}

func NewRoutineRouter(
	db *gorm.DB,
) *RoutineRouter {
	return &RoutineRouter{
		db:                db,
		routineRepository: repo.NewRoutineRepository(db),
	}
}

func (r *RoutineRouter) InitRoutineEndpoint(apiRouter *gin.RouterGroup) {

	apiRouter.GET("/routines", r.GetRoutines)
	apiRouter.GET("/routines/all", r.GetAllRoutines)
	apiRouter.POST("/routine", r.CreateRoutine)
	apiRouter.PUT("/routine/:id", r.UpdateRoutine)
	apiRouter.DELETE("/routine/:id", r.DeleteRoutine)
	apiRouter.POST("/routine/revert/:id", r.RevertDeleteRoutine)
}

func (r *RoutineRouter) CreateRoutine(c *gin.Context) {
	var routine model.Routine
	if err := c.Bind(&routine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* frequency must be at least everyday (1) */
	if routine.Frequency < 1 {
		routine.Frequency = 1
	}

	r.routineRepository.CreateRoutine(&routine)

	c.JSON(http.StatusCreated, routine)
}

func (r *RoutineRouter) GetRoutines(c *gin.Context) {
	var routines []model.Routine
	r.routineRepository.GetRoutines(&routines)

	c.JSON(http.StatusOK, routines)
}

func (r *RoutineRouter) GetAllRoutines(c *gin.Context) {
	var routines []model.Routine
	r.routineRepository.GetAllRoutines(&routines)

	c.JSON(http.StatusOK, routines)
}

func (r *RoutineRouter) UpdateRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var routine model.Routine
	if err := c.BindJSON(&routine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.routineRepository.UpdateRoutine(&routine, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, routine)
}

func (r *RoutineRouter) DeleteRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r.routineRepository.DeleteRoutine(&id)

	c.JSON(http.StatusOK, "Deleted successfully")
}

func (r *RoutineRouter) RevertDeleteRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = r.routineRepository.RevertDeleteRoutine(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "Reverted successfully")
}
