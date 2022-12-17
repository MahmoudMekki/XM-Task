package company

import (
	"github.com/MahmoudMekki/XM-Task/pkg/repo/companyDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func GetCompany(ctx *gin.Context) {
	company, err := companyDAL.GetCompany(ctx.GetString("id"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(company.Name) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no companies found!"})
		return
	}
	ctx.JSON(http.StatusOK, company)
}
