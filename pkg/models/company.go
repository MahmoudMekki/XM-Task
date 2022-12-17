package models

import "github.com/google/uuid"

const (
	CompaniesTableName = "companies"
)

type Company struct {
	Id           uuid.UUID `gorm:"column:id;varchar(16);primary_key" json:"id"`
	Name         string    `gorm:"column:name;varchar(20);not null;unique" json:"name"`
	Description  string    `gorm:"column:description;varchar()" json:"description"`
	EmployeesNum int64     `gorm:"column:employees_number; not null" json:"employees_number"`
	Registered   bool      `gorm:"column:registered; not null" json:"registered"`
	Type         string    `gorm:"column:type;type:enum('Corporations', 'NonProfit', 'Cooperative','Sole Proprietorship');not null" json:"type"`
}

func (c *Company) TableName() string {
	return CompaniesTableName
}
