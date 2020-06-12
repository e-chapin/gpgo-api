package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jackc/pgx/v4"
	"github.com/russross/blackfriday"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"


)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	defer db.Close()

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}



	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/mark", handleMark)
	router.GET("/items", dbHandler(db))

	router.Run(":" + port)
}

func dbHandler(db pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context){

		type Item struct {
			Id int
			Text string
			Description string
			Url string
			IsCompleted bool
		}

		var items []Item

		rows, err := db.Query(context.Background(), "select * from practice_item order by id asc")
		if err != nil 	{
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error Reading Practice Items"))
			return
		}

		//defer rows.Close()
		for rows.Next() {

			var i Item

			if err := rows.Scan(&i.Id, &i.Text, &i.Description, &i.Url, &i.IsCompleted); err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error scanning db: ", err))
			}
			items = append(items, i)
		}
		c.JSON(http.StatusOK, items)
	}
}



func handleMark( c *gin.Context) {
	c.String(http.StatusOK, string(blackfriday.Run([]byte("**Hi!**"))))
}
