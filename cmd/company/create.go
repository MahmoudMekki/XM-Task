package company

import (
	"encoding/json"
	"github.com/MahmoudMekki/XM-Task/pkg/kafka"
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
	companyobj, _ := json.Marshal(company)
	log := models.Log{
		UserId:   int64(ctx.GetFloat64("user_id")),
		Activity: models.CreateCompanyEvent,
		Data:     companyobj,
	}
	logObj, _ := json.Marshal(log)

	kafka.Produce(models.ActivityLogTopic, logObj)
	ctx.JSON(http.StatusCreated, company)
}
