package company

import (
	"encoding/json"
	"github.com/MahmoudMekki/XM-Task/pkg/kafka"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/pkg/repo/companyDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

func DeleteCompany(ctx *gin.Context) {
	err := companyDAL.DeleteCompany(ctx.GetString("id"))
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "company not found"})
			return
		}
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dataObj, _ := json.Marshal(gin.H{"id": ctx.GetString("id")})
	log := models.Log{
		UserId:   int64(ctx.GetFloat64("user_id")),
		Activity: models.DeleteCompanyEvent,
		Data:     dataObj,
	}
	logObj, _ := json.Marshal(log)

	kafka.Produce(models.ActivityLogTopic, logObj)
	ctx.JSON(http.StatusOK, gin.H{"id": ctx.GetString("id")})
}
