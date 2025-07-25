package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handler "zurihaqi.github.io-backend/internal/handler"
	middleware "zurihaqi.github.io-backend/internal/middleware"

	repository "zurihaqi.github.io-backend/internal/repository"
	service "zurihaqi.github.io-backend/internal/service"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)

	authService := service.NewAuthService(userRepo)
	projectService := service.NewProjectService(projectRepo)

	authHandler := handler.AuthHandler{Service: authService}
	projectHandler := handler.ProjectHandler{Service: projectService}

	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/projects", projectHandler.GetAllProjects)
		protected.GET("/projects/:id", projectHandler.GetProjectByID)
		protected.POST("/projects", projectHandler.CreateProject)
		protected.PUT("/projects/:id", projectHandler.UpdateProject)
		protected.DELETE("/projects/:id", projectHandler.DeleteProject)
	}
}
