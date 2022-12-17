package company

import (
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
	ctx.JSON(http.StatusOK, gin.H{"id": ctx.GetString("id")})
}
