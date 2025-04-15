package controllers

import (
	"errors"
	"net/http"

	"github.com/DevAthhh/DoZen/internal/auth"
	"github.com/DevAthhh/DoZen/internal/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(string, string, string) error
	GetUserByID(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
}

func NewLoginController(ur UserRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var bodyRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := ctx.Bind(&bodyRequest); err != nil {
			return err
		}

		user, _ := ur.GetUserByEmail(bodyRequest.Email)

		if user.ID == 0 {
			return errors.New("user doesn't exists")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyRequest.Password)); err != nil {
			return errors.New("password is incorrect")
		}

		token, err := auth.GenerateToken(int(user.ID), user.Email)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, struct {
			Msg    string `json:"token"`
			UserID uint   `json:"user_id"`
		}{Msg: token, UserID: user.ID})
	}
}

func NewRegisterController(ur UserRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var bodyRequest struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := ctx.Bind(&bodyRequest); err != nil {
			return err
		}

		if err := ur.CreateUser(bodyRequest.Username,
			bodyRequest.Email,
			bodyRequest.Password,
		); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, struct {
			Msg string `json:"msg"`
		}{Msg: "user has been created"})
	}
}

func NewProfileController(ur UserRepo) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")

		user, err := ur.GetUserByID(id)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	}
}
