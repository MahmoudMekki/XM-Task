package company

import (
	"fmt"
	conn "github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/tests/database"
	"github.com/MahmoudMekki/XM-Task/tests/server"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetCompanyHandler(t *testing.T) {
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

	t.Run("Get company successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/company/%s", company.Id), nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return bad request due to fake ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/company/%s", "blah12"), nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return not found due to fake ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/company/%s", "947638b9-2200-41d8-ab64-512002a921a0"), nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
