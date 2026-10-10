package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"coffee_ingredients/internal/app/handler"
	"coffee_ingredients/internal/app/repository"
)

func StartServer() {
	log.Println("Starting server")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./resources")

	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(repo)

	r.GET("/feed-coffee-ingredients", h.Feed)
	r.GET("/add-coffee-ingredients", h.Add)
	r.GET("/tiles-coffee-ingredients", h.Tiles)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/feed-coffee-ingredients")
	})

	r.Run()
}
