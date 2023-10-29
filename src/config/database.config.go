package config

import (
	"fmt"
	"go-gin-template/src/helper"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	log "github.com/sirupsen/logrus"
)

var DBConn *gorm.DB

func CreateDBConnection() {

	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database)

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Errorf("Error/Timeout occurred while connecting with the host %s", host)
		return
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Errorf("Cannot connect to DB (%s)", database)
		CloseDBConnection()
		return
	}

	idleConnection := helper.ParseStringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION"))
	maxOpenConnection := helper.ParseStringToInt(os.Getenv("DB_MAX_OPEN_CONNECTION"))

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(idleConnection)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenConnection)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	//uncomment to auto migrate
	// conn.AutoMigrate(
	// 	dao.Payment{}, dao.PaymentUniqueCode{}, dao.Currency{}, dao.CurrencyExchangeRate{},
	// 	dao.PaymentParty{}, dao.PaymentMethod{}, dao.CompanyCredential{}, 
	// dao.PaymentMethodTenant{}, dao.BankMapping{})
	DBConn = conn
	log.Printf("Success connect to DB %s", DBConn.Name())
}

func CloseDBConnection() {
	log.Info("Close DB connection")
	if DBConn != nil {
		sqlDB, err := DBConn.DB()
		if err != nil {
			log.Error("Error occurred while closing a DB connection")
		}
		defer sqlDB.Close()
	}
}