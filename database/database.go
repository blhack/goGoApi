package database

import (
	"database/sql"
	"fmt"
	"github.com/blhack/goGoApi/config"
	"github.com/blhack/goGoApi/utils"
)

var DBCon *sql.DB

func Init() {
	var err error
	DBCon, err = sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", config.DbUser, config.DbPass, config.DbName))
	utils.CheckErr(err)
}
