package company

import (
	"fmt"
	conn "github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/tests/database"
	"github.com/MahmoudMekki/XM-Task/tests/server"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDeleteCompanyHandler(t *testing.T) {
	os.Setenv("ENV", "test")
	database.DbSetup()
	db, err := conn.GetDatabaseConnection()
	if err != nil {
		t.Error(err)
	}
	defer database.CleanUpDb(db)
	r := server.SetUpRouter()
	company := models.Company{
		Id:           uuid.New(),
		Name:         "mekki",
		Description:  "blah",
		EmployeesNum: 500,
		Registered:   true,
		Type:         "NonProfit",
	}
	company = database.CreateCompany(db, company)
	t.Run("delete company successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/company/%s", company.Id), nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		_, err := database.GetCompany(db, company.Id.String())
		assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
	})
}
