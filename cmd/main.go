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
	brandService := service.NewBrandService(dao)
	countryService := service.NewCountryService(dao)
	designerService := service.NewDesignerService(dao)
	garmentService := service.NewGarmentService(dao)
	projectService := service.NewProjectService(dao)
	userService := service.NewUserService(dao)

	microService := app.NewMicroService(
		authService,
		brandService,
		countryService,
		designerService,
		garmentService,
		projectService,
		tokenManager,
		userService,
	)

	router := gin.Default()

	publicRoute := router.Group("/api/v1")

	userRoute := publicRoute.Group("/users")
	userRoute.POST("/", microService.CreateUser)
	userRoute.Use(microService.AuthorizeUser())
	userRoute.GET("/:id", microService.GetUser)
	userRoute.DELETE("/:id", microService.DeleteUser)
	userRoute.Use(microService.AuthorizeAdmin())
	userRoute.GET("/", microService.GetUsers)

	loginRoute := publicRoute.Group("/login")
	loginRoute.POST("/", microService.Login)

	countryRoute := publicRoute.Group("/countries")
	countryRoute.Use(microService.AuthorizeUser())
	countryRoute.GET("/:id", microService.GetCountry)
	countryRoute.Use(microService.AuthorizeAdmin())
	countryRoute.PUT("/:id", microService.UpdateCountry)
	countryRoute.POST("/", microService.CreateCountry)
	countryRoute.DELETE("/:id", microService.DeleteCountry)

	designerRoute := publicRoute.Group("/designers")
	designerRoute.Use(microService.AuthorizeUser())
	designerRoute.GET("/:id", microService.GetDesigner)
	designerRoute.Use(microService.AuthorizeAdmin())
	designerRoute.PUT("/:id", microService.UpdateDesigner)
	designerRoute.POST("/", microService.CreateDesigner)
	designerRoute.DELETE("/:id", microService.DeleteDesigner)

	brandRoute := publicRoute.Group("/brands")
	brandRoute.Use(microService.AuthorizeUser())
	brandRoute.GET("/:id", microService.GetBrand)
	brandRoute.Use(microService.AuthorizeAdmin())
	brandRoute.PUT("/:id", microService.UpdateBrand)
	brandRoute.POST("/", microService.CreateBrand)
	brandRoute.DELETE("/:id", microService.DeleteBrand)

	garmentRoute := publicRoute.Group("/garments")
	garmentRoute.Use(microService.AuthorizeUser())
	garmentRoute.GET("/:id", microService.GetGarment)
	garmentRoute.Use(microService.AuthorizeAdmin())
	garmentRoute.PUT("/:id", microService.UpdateGarment)
	garmentRoute.POST("/", microService.CreateGarment)
	garmentRoute.DELETE("/:id", microService.DeleteGarment)

	router.Run("localhost:8081")
}
