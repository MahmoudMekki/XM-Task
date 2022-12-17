package company

import (
	"bytes"
	"encoding/json"
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

func TestUpdateCompanyHandler(t *testing.T) {
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
	t.Run("update company successfully", func(t *testing.T) {
		payload := models.Company{
			Name:         "Mahmoud",
			EmployeesNum: 1000,
			Description:  "welcome to my company",
		}
		jsonVal, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/company/%s", company.Id), bytes.NewBuffer(jsonVal))
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		companyUpdated, err := database.GetCompany(db, company.Id.String())
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, payload.EmployeesNum, companyUpdated.EmployeesNum)
		assert.Equal(t, payload.Name, companyUpdated.Name)
		assert.Equal(t, payload.Description, companyUpdated.Description)
		assert.Equal(t, company.Type, companyUpdated.Type)
		assert.Equal(t, company.Registered, companyUpdated.Registered)
	})
}
