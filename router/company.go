package router

import (
	"github.com/MahmoudMekki/XM-Task/cmd/company"
	"github.com/MahmoudMekki/XM-Task/validators"
)

func (r *routerImp) setCompanyRoutes() {
	companyEndpoints := r.engine.Group("/company")
	companyEndpoints.POST("", validators.ValidateCreateCompany(), company.CreateCompany)
	companyEndpoints.GET("/:id", validators.ValidateGetDeleteUpdateCompany(), company.GetCompany)
	companyEndpoints.DELETE("/:id", validators.ValidateGetDeleteUpdateCompany(), company.DeleteCompany)
	companyEndpoints.PATCH("/:id", validators.ValidateGetDeleteUpdateCompany(), company.UpdateCompany)
}
