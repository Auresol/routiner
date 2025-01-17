package api

import (
	"routiner/server/src/api/router"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	routineRouter  *router.RoutineRouter
	taskRouter     *router.TaskRouter
	calenderRouter *router.CalenderRouter
	mockRouter     *router.MockRouter
	router         *gin.Engine
}

func NewApiRouter(
	routineRouter *router.RoutineRouter,
	taskRouter *router.TaskRouter,
	calenderRouter *router.CalenderRouter,
	mockRouter *router.MockRouter,
) *ApiRouter {
	router := gin.Default()
	return &ApiRouter{
		routineRouter:  routineRouter,
		taskRouter:     taskRouter,
		calenderRouter: calenderRouter,
		mockRouter:     mockRouter,
		router:         router,
	}
}

func (r *ApiRouter) InitRouter() {

	r.router.Use(CORSMiddleware())

	api := r.router.Group("/api")
	mock := r.router.Group("/mock")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.routineRouter.InitRoutineEndpoint(api)
	r.taskRouter.InitTaskEndpoint(api)
	r.calenderRouter.InitCalenderEndpoint(api)
	r.mockRouter.InitMockEndpoint(mock)
}

func (r *ApiRouter) Run(port string) {
	r.router.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
