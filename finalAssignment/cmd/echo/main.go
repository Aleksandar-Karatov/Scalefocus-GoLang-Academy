package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"final/cmd"
	"final/cmd/echo/logic"
	"final/cmd/echo/repository"
	"fmt"
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
	if err != nil {
		log.Fatal(err)
	}
	router.Use(middleware.BasicAuth(func(username, password string, _ echo.Context) (bool, error) {
		user, err := queries.GetUserByName(ctx, username)
		if err != nil {
			log.Println(err)
		}
		passwordHash := sha256.Sum256([]byte(password))
		if username == user.Username && fmt.Sprint(passwordHash) == user.Password {
			currentUserID = user.IDOfUser
			log.Println("current user is:", user.Username, " ", currentUserID)
			return true, nil

		}
		toInsert := repository.InsertUserInDBParams{}
		toInsert.Username = username
		toInsert.Password = fmt.Sprint(passwordHash)

		user, err = queries.InsertUserInDB(ctx, toInsert)
		if err != nil {
			log.Println(err)
		}
		currentUserID = user.IDOfUser
		log.Println("current user is:", user.Username, " ", currentUserID)
		return true, nil

	}))
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			asd := ctx.Request().URL
			log.Println("Echo middleware was called", asd)
			return next(ctx)
		}
	})

	router.GET("/api", logic.Logout(router.AcquireContext(), &currentUserID))
	router.GET("/api/lists", logic.GetLists(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.POST("/api/lists", logic.PostList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.GET("/api/lists/:id", logic.GetList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.DELETE("/api/lists/:id", logic.DeleteList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.GET("/api/lists/:id/tasks", logic.GetTasksForList(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.POST("/api/lists/:id/tasks", logic.PostTask(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.GET("/api/tasks/:id", logic.GetTask(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.DELETE("/api/tasks/:id", logic.DeleteTask(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.PATCH("/api/tasks/:id", logic.FinishTask(router.AcquireContext(), &currentUserID, queries, &ctx))
	router.GET("/api/weather", logic.GetWeather(router.AcquireContext()))
	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
