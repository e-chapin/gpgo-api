package controllers

import (
	"github.com/e-jameson/gpgo-api/src/db"
	"github.com/e-jameson/gpgo-api/src/helpers"
	"github.com/e-jameson/gpgo-api/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PracticeSessionController struct {}

func (ps PracticeSessionController) PracticeSession(c *gin.Context){

	pgdb := db.GetDB()

	id, err := helpers.GetPathInt(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "No Practice Session ID")
	}

	session := models.PracticeSession{Id: id}
	if err = pgdb.Select(&session); err != nil {
		c.JSON(http.StatusNotFound, "Practice Session Not Found")
	}

	c.JSON(http.StatusOK, session)
}

func (ps PracticeSessionController) AllPracticeSessions(c *gin.Context){

	pgdb := db.GetDB()

	var sessions []models.PracticeSession
	if err := pgdb.Model(&sessions).Order("id asc").Select(); err != nil {
		c.JSON(http.StatusBadRequest, "Error Loading Practice Sessions")
	}

	c.JSON(http.StatusOK, sessions)

}

func (ps PracticeSessionController) AddPracticeSession(c *gin.Context){

	pgdb := db.GetDB()

	var practiceSession models.PracticeSession
	if err := c.BindJSON(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error creating Practice Session")
	}
	if err := pgdb.Insert(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error creating Practice Session")
	}
	c.JSON(http.StatusOK, practiceSession)

}

func (ps PracticeSessionController) EditPracticeSession(c *gin.Context){

	pgdb := db.GetDB()

	var practiceSession models.PracticeSession
	if err := c.BindJSON(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error creating Practice Session")
	}
	if err := pgdb.Update(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error creating Practice Session")
	}
	if err := pgdb.Select(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error creating Practice Session")
	}
	c.JSON(http.StatusOK, practiceSession)

}

func (ps PracticeSessionController) DeletePracticeSession(c *gin.Context){

	pgdb := db.GetDB()

	var practiceSession models.PracticeSession
	if err := c.BindJSON(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error deleting Practice Session")
	}
	if err := pgdb.Delete(&practiceSession); err != nil {
		c.JSON(http.StatusBadRequest, "Error deleting Practice Session")
	}
	c.JSON(http.StatusOK, "OK")

}