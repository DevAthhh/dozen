package controllers

import (
	"net/http"

	"github.com/DevAthhh/DoZen/internal/models"
	"github.com/labstack/echo"
)

// TODO: изменить все значения id на int

type TaskRepo interface {
	CreateTask(groupID int, title string) error
	GetTasksByGroupID(id string) (*[]models.Task, error)
	UpdateTaskStatusByID(id, status string) error
	DeleteTaskByID(id string) error
}

type GroupRepo interface {
	CreateGroup(userIDs []int, name string) error
}

func NewCreateTaskController(tr TaskRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var bodyRequest struct {
			Title   string `json:"title"`
			GroupID int    `json:"group_id"`
		}

		if err := ctx.Bind(&bodyRequest); err != nil {
			return err
		}

		if err := tr.CreateTask(bodyRequest.GroupID, bodyRequest.Title); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, struct {
			Msg string `json:"msg"`
		}{Msg: "task has been created"})
	}
}

func NewUpdateTaskController(tr TaskRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var bodyRequest struct {
			Status string `json:"status"`
			ID     string `json:"id"`
		}

		if err := ctx.Bind(&bodyRequest); err != nil {
			return err
		}

		if err := tr.UpdateTaskStatusByID(bodyRequest.ID, bodyRequest.Status); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, struct {
			Msg string `json:"msg"`
		}{Msg: "task will be updated"})
	}
}

func NewDeleteTaskController(tr TaskRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var bodyRequest struct {
			ID string `json:"id"`
		}

		if err := ctx.Bind(&bodyRequest); err != nil {
			return err
		}

		if err := tr.DeleteTaskByID(bodyRequest.ID); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, struct {
			Msg string `json:"msg"`
		}{Msg: "task will be deleted"})
	}

}

func NewCreateGroupController(gr GroupRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var bodyRequest struct {
			Users []int  `json:"members"`
			Name  string `json:"name"`
		}

		if err := ctx.Bind(&bodyRequest); err != nil {
			return err
		}

		if err := gr.CreateGroup(bodyRequest.Users, bodyRequest.Name); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, struct {
			Msg string `json:"msg"`
		}{Msg: "group has been created"})
	}
}
