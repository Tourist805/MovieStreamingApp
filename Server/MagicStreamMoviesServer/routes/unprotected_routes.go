package routes

import (
	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUnProtectedRoutes(router *gin.Engine) {
	router.GET("/movies", controllers.GetMovies())
	router.GET("/genres", controllers.GetGenres())

	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
}
