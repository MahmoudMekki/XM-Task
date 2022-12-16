package models

type companyType string

const (
	CompaniesTableName             = "companies"
	Corporations       companyType = "Corporations"
	NonProfit          companyType = "NonProfit"
	Cooperative        companyType = "Cooperative"
	SoleProprietorship companyType = "Sole Proprietorship"
)

type Company struct {
	Id           string      `gorm:"column:id;type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string      `gorm:"column:name;varchar(20);not null;unique"`
	Description  string      `gorm:"column:description;varchar(3000); null"`
	EmployeesNum int64       `gorm:"column:employees_number; not null"`
	Type         companyType `gorm:"column:type;type:enum('Corporations', 'NonProfit', 'Cooperative','Sole Proprietorship');not null"`
}

func (c *Company) TableName() string {
	return CompaniesTableName
}
