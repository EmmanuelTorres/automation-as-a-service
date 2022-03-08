package app

import "github.com/automation-as-a-service/internal/service"

type MicroserviceServer struct {
	authService     service.AuthService
	brandService    service.BrandService
	countryService  service.CountryService
	designerService service.DesignerService
	garmentService  service.GarmentService
	projectService  service.ProjectService
	tokenManager    service.TokenManager
	userService     service.UserService
}

func NewMicroService(
	authService service.AuthService,
	brandService service.BrandService,
	countryService service.CountryService,
	designerService service.DesignerService,
	garmentService service.GarmentService,
	projectService service.ProjectService,
	tokenManager service.TokenManager,
	userService service.UserService,
) *MicroserviceServer {
	return &MicroserviceServer{
		authService:     authService,
		brandService:    brandService,
		countryService:  countryService,
		designerService: designerService,
		garmentService:  garmentService,
		projectService:  projectService,
		tokenManager:    tokenManager,
		userService:     userService,
	}
}
