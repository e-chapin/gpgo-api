package main

import (
	"flag"
	"fmt"
	"github.com/e-jameson/gpgo/config"
	"github.com/e-jameson/gpgo/db"
	"github.com/e-jameson/gpgo/server"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	//"github.com/go-pg/pg/v10/orm"

)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}