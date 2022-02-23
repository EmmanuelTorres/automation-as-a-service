package app

import "github.com/automation-as-a-service/internal/service"

type MicroserviceServer struct {
	authService    service.AuthService
	countryService service.CountryService
	tokenManager   service.TokenManager
	userService    service.UserService
}

func NewMicroService(
	authService service.AuthService,
	countryService service.CountryService,
	tokenManager service.TokenManager,
	userService service.UserService,
) *MicroserviceServer {
	return &MicroserviceServer{
		authService:    authService,
		countryService: countryService,
		tokenManager:   tokenManager,
		userService:    userService,
	}
}
