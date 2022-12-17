package router

import (
	"github.com/MahmoudMekki/XM-Task/cmd/authentication"
	"github.com/MahmoudMekki/XM-Task/cmd/company"
	"github.com/MahmoudMekki/XM-Task/validators"
)

func (r *routerImp) setCompanyRoutes() {
	companyEndpoints := r.engine.Group("/company")
	companyEndpoints.POST("", authentication.IsAuthorized(), validators.ValidateCreateCompany(), company.CreateCompany)
	companyEndpoints.GET("/:id", validators.ValidateGetDeleteUpdateCompany(), company.GetCompany)
	companyEndpoints.DELETE("/:id", authentication.IsAuthorized(), validators.ValidateGetDeleteUpdateCompany(), company.DeleteCompany)
	companyEndpoints.PATCH("/:id", authentication.IsAuthorized(), validators.ValidateGetDeleteUpdateCompany(), company.UpdateCompany)
}
