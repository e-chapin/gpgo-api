package server

import (
	"github.com/e-jameson/gpgo/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GPRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")


	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Simple group: v1
	item := router.Group("/item")
	{
		itemsController := new(controllers.ItemController)
		item.GET("/id/:id", itemsController.Item)
		item.GET("/all", itemsController.AllItems)
	}

	return router
}
