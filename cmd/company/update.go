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

func UpdateCompany(ctx *gin.Context) {
	var company models.Company
	err := ctx.ShouldBindJSON(&company)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company.Id, _ = uuid.Parse(ctx.GetString("id"))
	err = companyDAL.UpdateCompany(company)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companyobj, _ := json.Marshal(company)
	log := models.Log{
		UserId:   int64(ctx.GetFloat64("user_id")),
		Activity: models.UpdateCompanyEvent,
		Data:     companyobj,
	}
	logObj, _ := json.Marshal(log)

	kafka.Produce(models.ActivityLogTopic, logObj)
	ctx.JSON(http.StatusOK, gin.H{"id": ctx.GetString("id")})
}
