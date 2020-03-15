package dbaccess

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ConnectArgs struct {
	Address  string
	Port     string
	DBName   string
	User     string
	Password string
}

func (conargs *ConnectArgs) SetDefault() {
	connectTemplate := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	defaultConnectString = fmt.Sprintf(connectTemplate, conargs.User, conargs.Password, conargs.Address, conargs.Port, conargs.DBName)
}

var defaultConnectString string

func ConnectGorm() *gorm.DB {
	log.Println(defaultConnectString)
	db, err := gorm.Open("mysql", defaultConnectString)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
