package company

import (
	"bytes"
	"encoding/json"
	"fmt"
	conn "github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/tests/database"
	"github.com/MahmoudMekki/XM-Task/tests/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateCompanyHandler(t *testing.T) {
	os.Setenv("ENV", "test")
	database.DbSetup()
	db, err := conn.GetDatabaseConnection()
	if err != nil {
		t.Error(err)
	}
	defer database.CleanUpDb(db)
	r := server.SetUpRouter()
	company := models.Company{
		Name:         "mekki",
		Description:  "blah",
		EmployeesNum: 500,
		Registered:   true,
		Type:         "NonProfit",
	}
	t.Run("create company successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		jsonVal, _ := json.Marshal(company)
		req, _ := http.NewRequest("POST", fmt.Sprintf("/company"), bytes.NewBuffer(jsonVal))
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("should return bad request due to empty name", func(t *testing.T) {
		w := httptest.NewRecorder()
		company.Type = "ay 7aga"
		jsonVal, _ := json.Marshal(company)
		req, _ := http.NewRequest("POST", fmt.Sprintf("/company"), bytes.NewBuffer(jsonVal))
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return bad request due to company name is already existed", func(t *testing.T) {
		w := httptest.NewRecorder()
		jsonVal, _ := json.Marshal(company)
		req, _ := http.NewRequest("POST", fmt.Sprintf("/company"), bytes.NewBuffer(jsonVal))
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
