package handler

import (
	"fmt"

	"github.com/DevAthhh/DoZen/internal/controllers"
	"github.com/DevAthhh/DoZen/internal/middlewares"
	"github.com/DevAthhh/DoZen/internal/repository"
	"github.com/DevAthhh/DoZen/pkg/lib/config"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type repo interface {
	NewTaskRepo() *repository.TaskRepository
	NewUserRepo() *repository.UserRepository
	NewGroupRepo() *repository.GroupRepository
}

func Route(logger *zap.Logger, repo repo) *echo.Echo {
	router := echo.New()
	version := config.Cfg.APIVersion

	router.Use(middlewares.LoadLoggerMiddleware(logger))

	api := router.Group("/api")
	{
		ver := api.Group(fmt.Sprintf("/%s", version))
		{
			auth := ver.Group("/u")
			{
				auth.POST("/login", controllers.NewLoginController(repo.NewUserRepo()))
				auth.POST("/register", controllers.NewRegisterController(repo.NewUserRepo()))
				auth.GET("/:id", controllers.NewProfileController(repo.NewUserRepo()))
			}

			task := ver.Group("/t")
			{
				task.POST("", controllers.NewCreateTaskController(repo.NewTaskRepo()))
				task.PUT("", controllers.NewUpdateTaskController(repo.NewTaskRepo()))
				task.DELETE("", controllers.NewDeleteTaskController(repo.NewTaskRepo()))
			}

			group := ver.Group("/g")
			{
				group.POST("", controllers.NewCreateGroupController(repo.NewGroupRepo()))
			}
		}
	}

	return router
}
