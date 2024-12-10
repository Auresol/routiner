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

	fmt.Print("Currently on " + env + " enviroment")

	router := setup(env)
	router.InitRouter()
	router.Run(":8080")
}

func setup(env string) *api.ApiRouter {

	db := db.NewDBConnection(env)

	db.AutoMigrate(model.Task{})
	db.AutoMigrate(model.Log{})
	db.AutoMigrate(model.Calender{})

	taskRouter := router.NewTaskRouter(db)
	logRouter := router.NewLogRouter(db)
	calenderRouter := router.NewCalenderRouter(db)
	mockRouter := router.NewMockRouter(db)

	router := api.NewApiRouter(
		taskRouter,
		logRouter,
		calenderRouter,
		mockRouter,
	)

	return router

}
