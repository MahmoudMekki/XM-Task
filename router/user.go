package router

import (
	"github.com/MahmoudMekki/XM-Task/cmd/users"
	"github.com/MahmoudMekki/XM-Task/validators"
)

func (r *routerImp) setUserRoutes() {
	userEndpoints := r.engine.Group("/user")
	userEndpoints.POST("/signup", validators.ValidateSignup(), users.Signup)
	userEndpoints.GET("/login", validators.ValidateLogin(), users.Login)
}
