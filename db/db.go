package db

import (
	"github.com/e-jameson/gpgo/config"
	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func Init() {

	c := config.GetConfig()

	opt, err := pg.ParseURL(c.GetString("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)
}


func GetDB() *pg.DB {
	return db
}

//func createSchema(db *pg.DB) error {
//	models := []interface{}{
//		(*PracticeItem)(nil),
//		(*Item)(nil),
//	}
//
//	for _, model := range models {
//		err := db.CreateTable(model, &orm.CreateTableOptions{
//			Temp: true, // temp table
//		})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

