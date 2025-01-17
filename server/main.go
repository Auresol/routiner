package main

import (
	"fmt"
	"os"
	"routiner/server/src/api"
	"routiner/server/src/api/router"
	"routiner/server/src/db"
	"routiner/server/src/model"
)

func main() {

	env := os.Args[1]

	if env == "" {
		env = "dev"
	}

	fmt.Println("Currently on " + env + " enviroment")

	router := setup(env)
	router.InitRouter()
	router.Run(":8080")
}

func setup(env string) *api.ApiRouter {

	db := db.NewDBConnection(env)

	db.AutoMigrate(model.Routine{})
	db.AutoMigrate(model.Task{})
	db.AutoMigrate(model.Calender{})

	calenderRouter := router.NewCalenderRouter(db)
	routineRouter := router.NewRoutineRouter(db)
	taskRouter := router.NewTaskRouter(db)
	mockRouter := router.NewMockRouter(db)

	router := api.NewApiRouter(
		routineRouter,
		taskRouter,
		calenderRouter,
		mockRouter,
	)

	return router

}
