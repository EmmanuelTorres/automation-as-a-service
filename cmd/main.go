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

	// Postgres
	db, err := repository.NewDB()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	// Prepare config file
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

	publicRoute := router.Group("/api/v1")

	userRoute := publicRoute.Group("/users")
	userRoute.Use(microService.AuthorizeUser())
	userRoute.GET("/:id", microService.GetUser)
	userRoute.POST("/", microService.CreateUser)
	userRoute.DELETE("/:id", microService.DeleteUser)

	loginRoute := publicRoute.Group("/login")
	loginRoute.POST("/", microService.Login)

	projectRoute := publicRoute.Group("/projects")
	projectRoute.GET("/:name", microService.GetProject)
	projectRoute.POST("", microService.CreateProject)
	projectRoute.PUT("/:id", microService.UpdateProject)
	projectRoute.DELETE("/:id", microService.DeleteProject)

	// router.Use(microService.AuthorizeUser())

	router.Run("localhost:8081")
}
