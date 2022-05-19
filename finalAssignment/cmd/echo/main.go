package main

import (
	"context"
	"database/sql"
	"final/cmd"
	"final/cmd/echo/logicForApp"
	"final/cmd/echo/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

func main() {
	var currentUserID int64
	currentUserID = -1
	router := echo.New()
	db, err := sql.Open("sqlite", "DBFinal.db")
	ctx := context.Background()
	queries := repository.New(db)
	//api:= router.Group("/api")
	if err != nil {
		log.Fatal(err)
	}
	router.Use(middleware.BasicAuth(func(username, password string, _ echo.Context) (bool, error) {
		//log.Println(db.Query("SELECT * FROM users"))
		user, err := queries.GetUserByName(ctx, username)
		if err != nil {
			log.Println(err)
		}
		if username == user.Username && password == user.Password {
			currentUserID = user.IDOfUser
			log.Println("current user is:", user.Username, " ", currentUserID)
			return true, nil

		}
		return false, nil

	}))
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			asd := ctx.Request().URL
			log.Println("Echo middleware was called", asd)
			return next(ctx)
		}
	})

	router.GET("/api", logicForApp.Logout(router.AcquireContext(), &currentUserID))
	router.GET("/api/lists", logicForApp.GetLists(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.POST("/api/lists", logicForApp.PostList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.GET("/api/lists/:id", logicForApp.GetList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.DELETE("/api/lists/:id", logicForApp.DeleteList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.GET("/api/lists/:id/tasks", logicForApp.GetTasksForList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.POST("/api/lists/:id/tasks", logicForApp.PostTask(router.AcquireContext(), &currentUserID, queries, &ctx))
	//router.DELETE("/api/lists/:id/tasks", )
	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
