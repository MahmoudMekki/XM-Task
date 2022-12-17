package validators

import (
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

func ValidateCreateCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		data, err := jio.ValidateJSON(&jsonData, jio.Object().Keys(jio.K{
			"name":             jio.String().Max(15).Required(),
			"description":      jio.String().Max(3000).Optional(),
			"employees_number": jio.Number().Required(),
			"registered":       jio.Bool().Required(),
			"type":             jio.String().Valid("Corporations", "NonProfit", "Cooperative", "Sole Proprietorship").Required(),
		}))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("name", data["name"])
		ctx.Set("description", data["description"])
		ctx.Set("employees_number", data["employees_number"])
		ctx.Set("registered", data["registered"])
		ctx.Set("type", data["type"])
		ctx.Next()
	}
}

func ValidateGetDeleteUpdateCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id is missed"})
			return
		}
		_, err := uuid.Parse(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id isn't valid"})
			return
		}
		ctx.Set("id", id)
		ctx.Next()
	}
}
