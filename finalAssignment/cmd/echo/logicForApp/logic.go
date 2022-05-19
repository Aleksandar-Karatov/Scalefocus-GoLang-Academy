package logicForApp

import (
	"context"
	"encoding/json"
	"final/cmd/echo/repository"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Logout(ctx echo.Context, curUser *int64) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		*curUser = -1
		return ctx.JSON(401, "LOGOUT")
	}
}

//func(username, password string, c echo.Context)
// type GetListsPagePayload struct {
// 	AllListsForUser []repository.GetListsForCurrentUserRow `json:"all"`
// }

func GetLists(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		allListsForCurUser, err := queries.GetListsForCurrentUser(*dbContext, *curUser)
		if err != nil {
			log.Println(err)
		}
		if len(allListsForCurUser) == 0 {
			toAdd := repository.InsertListInDBParams{Name: "sample task", Userid: *curUser}
			queries.InsertListInDB(*dbContext, toAdd)
			allListsForCurUser, err = queries.GetListsForCurrentUser(*dbContext, *curUser)
			if err != nil {
				log.Println(err)
			}
		}
		return ctx.JSON(200, allListsForCurUser)
	}
}

type listToPost struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func PostList(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		toAdd := repository.InsertListInDBParams{}
		json.NewDecoder(ctx.Request().Body).Decode(&toAdd)
		toAdd.Userid = *curUser
		log.Println(toAdd)
		added, err := queries.InsertListInDB(*dbContext, toAdd)
		if err != nil {
			log.Println(err)
		}
		log.Println("Added list in db:", added.Name)
		toPost := listToPost{Id: added.IDOfList, Name: added.Name}
		return ctx.JSON(200, toPost)
	}
}

func DeleteList(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Println("ctx param", ctx.Param("id"))
		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		log.Println("ID to delete is:", id)
		allTasks, err := queries.GetTasksForCurrentList(*dbContext, id)
		if err != nil {
			log.Println(err)
		}
		for _, task := range allTasks {
			_ = queries.DeleteTasktByID(*dbContext, task.IDOfTask)
		}

		data := queries.DeleteListByID(*dbContext, id)
		return ctx.JSON(200, data)
	}

}
func GetList(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Println("Called get list")
		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

		toGetAllLists, err := queries.GetListsForCurrentUser(*dbContext, *curUser)
		if err != nil {
			log.Println(err)
		}
		var onlyListToGet repository.GetListsForCurrentUserRow
		for _, item := range toGetAllLists {
			if item.IDOfList == id {
				onlyListToGet.IDOfList = item.IDOfList
				onlyListToGet.Name = item.Name
			}
		}
		return ctx.JSON(200, onlyListToGet)
	}

}
func GetTasksForList(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		toGetAllTasks, err := queries.GetTasksForCurrentList(*dbContext, id)
		if err != nil {
			log.Println(err)
		}
		if len(toGetAllTasks) == 0 {
			toGetAllTasks = append(toGetAllTasks, repository.Task{})
			return ctx.JSON(200, toGetAllTasks)
		}
		return ctx.JSON(200, toGetAllTasks)
	}
}

type taskToPost struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	ListId    int64  `json:"listId"`
	Completed bool   `json:"completed"`
}

func PostTask(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		toAdd := repository.InsertTaskInDBParams{}
		json.NewDecoder(ctx.Request().Body).Decode(&toAdd)
		idOfList, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		toAdd.Listid = idOfList
		toAdd.Completed = false
		log.Println(toAdd)
		added, err := queries.InsertTaskInDB(*dbContext, toAdd)
		if err != nil {
			log.Println(err)
		}
		log.Println("Added list in db:", added.Text)
		toPost := taskToPost{Id: added.IDOfTask, Text: added.Text, ListId: toAdd.Listid, Completed: added.Completed}
		return ctx.JSON(200, toPost)
	}
}

func DeleteTask(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Println("ctx param", ctx.Param("id"))
		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		log.Println("ID to delete is:", id)

		data := queries.DeleteTasktByID(*dbContext, id)
		return ctx.JSON(200, data)

	}

}
func GetTask(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Println("Called get task")
		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

		getTask, err := queries.GetTaskByID(*dbContext, id)
		if err != nil {
			log.Println(err)
		}

		return ctx.JSON(200, getTask)
	}
}

func FinishTask(ctx echo.Context, curUser *int64, queries *repository.Queries, dbContext *context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Println("finish task is called")
		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		toUpdate := repository.PatchTaskInDBParams{}
		toUpdate.IDOfTask = id
		json.NewDecoder(ctx.Request().Body).Decode(&toUpdate)
		err := queries.PatchTaskInDB(*dbContext, toUpdate)
		if err != nil {
			log.Println(err)
			return ctx.JSON(404, nil)
		}
		patchedData, err := queries.GetTaskByID(*dbContext, id)
		if err != nil {
			log.Println(err)
		}

		return ctx.JSON(200, patchedData)
	}

}