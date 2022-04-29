package routes

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/")
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
	}

	userLoginRoutes := r.Group("/users")
	{
		userLoginRoutes.Use(middlewares.Authentication())

		userLoginRoutes.PUT("/:userId", controllers.UpdateUser)
		userLoginRoutes.DELETE("/:userId", controllers.DeleteUser)
	}

	photoRoutes := r.Group("/photos")
	{
		photoRoutes.Use(middlewares.Authentication())

		photoRoutes.POST("/", controllers.NewPhoto)
		photoRoutes.GET("/", controllers.GetPhotos)
		photoRoutes.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRoutes := r.Group("/comments")
	{
		commentRoutes.Use(middlewares.Authentication())

		commentRoutes.POST("/", middlewares.CommentCreateAuthorization(), controllers.NewComment)
		commentRoutes.GET("/", controllers.GetComments)
		commentRoutes.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRoutes.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRoutes := r.Group("/socialmedias")
	{
		socialMediaRoutes.Use(middlewares.Authentication())

		socialMediaRoutes.POST("/", controllers.NewSocialMedia)
		socialMediaRoutes.GET("/", controllers.GetSocialMedia)
		socialMediaRoutes.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRoutes.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
