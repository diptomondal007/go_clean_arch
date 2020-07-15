package http

import (
	"github.com/diptomondal007/go_clean_arch/auth"
	"github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Echo, useCase auth.UseCase){
	h := NewHandler(useCase)

	authEndpoints := router.Group("/auth")
	authEndpoints.POST("/sign-up",h.SignUp)
	authEndpoints.POST("/sign-in", h.SignIn)

}
