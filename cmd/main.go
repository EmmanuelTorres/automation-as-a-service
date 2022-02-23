package main

import (
	"log"

	"github.com/automation-as-a-service/internal/app"
	"github.com/automation-as-a-service/internal/repository"
	"github.com/automation-as-a-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// DB
	db, err := repository.NewDB()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	// preparing config file
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}

	// JWT
	signedKeyJwt := viper.Get("jwt.signedKey").(string)
	tokenManager := service.NewTokenManager(signedKeyJwt)

	// Register all services
	dao := repository.NewDAO(db)
	authService := service.NewAuthService(dao, tokenManager)
	countryService := service.NewCountryService(dao)
	projectService := service.NewProjectService(dao)
	userService := service.NewUserService(dao)

	microService := app.NewMicroService(
		authService,
		countryService,
		projectService,
		tokenManager,
		userService,
	)

	router := gin.Default()

	// User
	router.GET("/user/:id", microService.GetUser)
	router.POST("/user", microService.CreateUser)
	router.DELETE("/user/:id", microService.DeleteUser)
	router.POST("/login", microService.Login)

	// Country
	router.GET("/country/:name", microService.GetCountry)
	router.POST("/country", microService.CreateCountry)
	router.DELETE("/country/:name", microService.DeleteCountry)

	// Project
	router.POST("/project", microService.CreateProject)

	router.Run("localhost:8081")
}
