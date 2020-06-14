package controllers

import (
	"github.com/e-jameson/gpgo/db"
	"github.com/e-jameson/gpgo/helpers"
	"github.com/e-jameson/gpgo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemController struct {}

func (i ItemController) Item (c *gin.Context){

	pgdb := db.GetDB()

	id, err := helpers.GetPathInt(c, "id")
	if err != nil {
		panic(err)
	}

	item := models.PracticeItem{Id: id}
	err = pgdb.Select(&item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, item)
}

func (i ItemController) AllItems (c *gin.Context) {

	pgdb := db.GetDB()

	// Select all users.
	var pi []models.PracticeItem
	err := pgdb.Model(&pi).Select()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, pi)
}

func (i ItemController) AddItem (c *gin.Context) {

	pgdb := db.GetDB()

	var item models.PracticeItem
	err := c.BindJSON(&item)
	if err != nil{
		panic(err)
	}

	err = pgdb.Insert(&item)
	if err != nil{
		panic(err)
	}
	c.JSON(http.StatusOK, item)
}


func (i ItemController) EditItem (c *gin.Context) {

	pgdb := db.GetDB()

	var item models.PracticeItem
	err := c.BindJSON(&item)
	if err != nil{
		panic(err)
	}

	err = pgdb.Update(&item)
	if err != nil{
		panic(err)
	}
	err = pgdb.Select(&item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, item)
}

func (i ItemController) DeleteItem (c *gin.Context){

	pgdb := db.GetDB()

	id, err := helpers.GetPathInt(c, "id")

	item := models.PracticeItem{Id: id}

	err = pgdb.Delete(&item)
	if err != nil{
		panic(err)
	}

	c.JSON(http.StatusOK, "OK")

}