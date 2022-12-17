package users

import (
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/pkg/repo/userDAL"
	"github.com/MahmoudMekki/XM-Task/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

func Signup(ctx *gin.Context) {
	user := models.User{
		UserName: ctx.GetString("user_name"),
		Email:    ctx.GetString("email"),
		Password: utils.Hash([]byte(ctx.GetString("password"))),
	}
	exists, err := userDAL.IsEmailExists(user.Email)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}
	exists, err = userDAL.IsUserNameExists(user.UserName)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user name already exists"})
		return
	}
	user, err = userDAL.CreateUser(user)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}
