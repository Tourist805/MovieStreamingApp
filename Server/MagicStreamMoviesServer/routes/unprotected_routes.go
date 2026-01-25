package routes

import (
	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.GET("/movies", handlers.GetMovies(client))
	router.GET("/genres", handlers.GetGenres(client))

	router.POST("/register", handlers.RegisterUser(client))
	router.POST("/login", handlers.LoginUser(client))
}
