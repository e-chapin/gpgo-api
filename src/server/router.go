package server

import (
	"fmt"
	"github.com/e-jameson/gpgo-api/src/config"
	"github.com/e-jameson/gpgo-api/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GPRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	router.LoadHTMLGlob("src/static/*.html")
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

	sessions := router.Group("/session")
	{
		sessionController := new(controllers.PracticeSessionController)
		sessions.GET("/id/:id", sessionController.PracticeSession)
		sessions.GET("/all", sessionController.AllPracticeSessions)
		sessions.DELETE("/all/:id", sessionController.AllPracticeSessions)
		sessions.POST("/new", sessionController.AllPracticeSessions)
		sessions.POST("/edit", sessionController.AllPracticeSessions)
	}

	return router
}
