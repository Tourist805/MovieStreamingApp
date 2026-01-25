package routes

import (
	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/handlers"
	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", handlers.GetMovie(client))
	router.POST("/addmovie", handlers.AddMovie(client))
	router.GET("/recommendedmovies", handlers.GetRecommendedMovies(client))
	router.PATCH("/updatereview/:imdb_id", handlers.AdminReviewUpdate(client))
}
