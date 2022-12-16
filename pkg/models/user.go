package models

import "time"

const UsersTableName = "users"

type User struct {
	Id        uint64    `gorm:"column:id; primary_key;auto_increment"`
	UserName  string    `gorm:"column:user_name; varchar(50); not null;unique"`
	Email     string    `gorm:"column: email; varchar(50);not null; unique"`
	Password  string    `gorm:"column:password;varchar(50);not null"`
	CreatedAt time.Time `gorm:"column: created_at; default: CURRENT_TIMESTAMP;not null"`
	UpdatedAt time.Time `gorm:"column: updated_at; default: CURRENT_TIMESTAMP;not null"`
}

func (u *User) TableName() string {
	return UsersTableName
}
