package schemas

import (
	"github.com/gookit/validate"
)

type Employee struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement" validate:"required|id"`
	LoginId   string `json:"login_id" gorm:"name" validate:"required|login_id"`
	FirstName string `json:"first_name" gorm:"first_name" validate:"required|first_name"`
	LastName  string `json:"last_name" gorm:"last_name" validate:"required|last_name"`
	UserName  string `json:"user_name" gorm:"user_name" validate:"required|user_name"`
}

// Messages you can custom validator error messages.
func (e Employee) Messages() map[string]string {
	return validate.MS{
		"required": "oh! the {field} is required",
		"id":       "Invalid id format",
		"name":     "Invalid name format",
	}
}

func (Employee) TableName() string {
	return "employee"
}
