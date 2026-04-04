package router

import (
	"github.com/JoaoGSantiago/starti-backend/internal/handlers"
	"github.com/JoaoGSantiago/starti-backend/internal/middleware"
	"github.com/JoaoGSantiago/starti-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func Setup(
	jwtService services.JWTService,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	postHandler *handlers.PostHandler,
	commentHandler *handlers.CommentHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	// rotas públicas
	api.POST("/auth/login", authHandler.Login)
	api.POST("/users", userHandler.CreateUser)

	// rotas protegidas
	protected := api.Group("")
	protected.Use(middleware.Auth(jwtService))
	{
		users := protected.Group("/users")
		users.GET("", userHandler.ListUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/:id/posts", userHandler.ListUserPosts)
		users.GET("/:id/comments", userHandler.ListUserComments)

		posts := protected.Group("/posts")
		posts.POST("", postHandler.CreatePost)
		posts.GET("/:id", postHandler.GetPost)
		posts.PUT("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
		posts.PATCH("/:id/archive", postHandler.ArchivePost)
		posts.GET("/:id/comments", postHandler.ListPostComments)

		comments := protected.Group("/comments")
		comments.POST("", commentHandler.CreateComment)
		comments.PUT("/:id", commentHandler.UpdateComment)
		comments.DELETE("/:id", commentHandler.DeleteComment)
	}

	return r
}
