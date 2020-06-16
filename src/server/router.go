package server

import (
	"github.com/e-jameson/gpgo/src/config"
	"github.com/e-jameson/gpgo/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GPRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("static/*.html")
	router.Static("/static", "static")


	c := config.GetConfig()
	routerConfig := cors.DefaultConfig()

	s := strings.Split(c.GetString("CORS"), ", ")

	routerConfig.AllowOrigins = s

	router.Use(cors.New(routerConfig))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Simple group: v1
	item := router.Group("/item")
	{
		itemsController := new(controllers.ItemController)
		item.GET("/id/:id", itemsController.Item)
		item.GET("/all", itemsController.AllItems)
		item.DELETE("/id/:id", itemsController.DeleteItem)
		item.POST("/new", itemsController.AddItem)
		item.POST("/edit", itemsController.EditItem)
		//item.POST("/complete", itemsController.CompleteItem)
	}

	return router
}
