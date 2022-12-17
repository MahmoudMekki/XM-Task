package company

import (
	"github.com/MahmoudMekki/XM-Task/pkg/repo/companyDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func DeleteCompany(ctx *gin.Context) {
	err := companyDAL.DeleteCompany(ctx.GetString("id"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": ctx.GetString("id") })
}