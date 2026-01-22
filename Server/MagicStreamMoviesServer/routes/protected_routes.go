package routes

import (
	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/controllers"
	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())
	router.GET("/recommendedmovies", controllers.GetRecommendedMovies())
}
