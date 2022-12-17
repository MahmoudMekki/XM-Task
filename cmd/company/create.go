package company

import (
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/pkg/repo/companyDAL"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateCompany(ctx *gin.Context) {
	company := models.Company{
		Id:           uuid.New(),
		Name:         ctx.GetString("name"),
		Description:  ctx.GetString("description"),
		EmployeesNum: int64(ctx.GetFloat64("employees_number")),
		Type:         ctx.GetString("type"),
		Registered:   ctx.GetBool("registered"),
	}
	company, err := companyDAL.CreateCompany(company)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, company)
}
