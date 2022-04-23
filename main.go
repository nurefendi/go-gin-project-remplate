package main

import (
	"crowdfounding/handler"
	"crowdfounding/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// use godot package to load/read the .env file and
// return the value of the key
func env(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }
  

func main() {
	dsn := env("DB_USER")+":"+env("DB_PASS")+"@tcp("+env("DB_HOST")+":"+env("DB_PORT")+")/"+env("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group(env("path"))

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	router.Run(env("host")+":"+env("port"))

}
