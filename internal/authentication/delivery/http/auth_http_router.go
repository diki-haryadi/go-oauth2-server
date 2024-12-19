package authHttpController

import (
	"github.com/labstack/echo/v4"

	oauthDomain "github.com/diki-haryadi/go-micro-template/internal/authentication/domain"
)

type Router struct {
	controller oauthDomain.HttpController
}

func NewRouter(controller oauthDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	auth := e.Group("/auth")
	{
		auth.POST("/register", r.controller.Register)
		auth.POST("/change-password", r.controller.ChangePassword)
		auth.POST("/forgot-password", r.controller.ForgotPassword)
		auth.POST("/update-username", r.controller.UpdateUsername)
	}

}
