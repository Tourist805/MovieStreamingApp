package main

import (
	"fmt"

	"github.com/Tourist805/MovieStreamingApp/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/genres", controllers.GetGenres())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())

	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
