package server

import (
	"github.com/e-jameson/gpgo-api/src/config"
)


func Init() {

	router := GPRouter()

	c := config.GetConfig()
	c.Get("port")

	router.Run(":" + c.GetString("Port"))

}