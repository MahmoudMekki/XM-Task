package models

const UsersTableName = "users"

type User struct {
	Id       int64  `gorm:"column:id; primary_key;auto_increment" json:"id"`
	UserName string `gorm:"column:user_name; varchar(50); not null;unique" json:"user_name"`
	Email    string `gorm:"column:email; varchar(50);not null; unique" json:"email"`
	Password string `gorm:"column:password;varchar(50);not null" json:"-"`
}

func (u *User) TableName() string {
	return UsersTableName
}
