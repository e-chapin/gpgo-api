package db

import (
	"github.com/e-jameson/gpgo-api/src/config"
	"github.com/e-jameson/gpgo-api/src/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func Init() {

	c := config.GetConfig()

	opt, err := pg.ParseURL(c.GetString("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)


	// create tables
	//todo where does this go
	models := []interface{}{
		(*models.PracticeItem)(nil),
		(*models.PracticeSession)(nil),
		(*models.SessionItem)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		}
	}


}

func GetDB() *pg.DB {
	return db
}

