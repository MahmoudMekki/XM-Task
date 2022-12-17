package users

import (
	"github.com/MahmoudMekki/XM-Task/pkg/repo/userDAL"
	"github.com/MahmoudMekki/XM-Task/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(ctx *gin.Context) {
	user, err := userDAL.GetUserByEmail(ctx.GetString("email"))
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			ctx.JSON(400, gin.H{"error": "invalid credentials"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
	valid := utils.ComparePasswordHashed(user.Password, []byte(ctx.GetString("password")))
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}
	token := utils.GenerateToken(user.Id)
	ctx.JSON(http.StatusOK, gin.H{"user_info": user, "token": token})
}
