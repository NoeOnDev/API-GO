package models

import (
    "github.com/go-playground/validator/v10"
)

type User struct {
    TableName struct{} `pg:"users"`
    ID        int      `pg:"id,pk"`
    FirstName string   `pg:"first_name" validate:"required"`
    LastName  string   `pg:"last_name" validate:"required"`
    Username  string   `pg:"username,unique" validate:"required"`
    Phone     string   `pg:"phone" validate:"required"`
    Email     string   `pg:"email,unique" validate:"required,email"`
    Password  string   `pg:"password" validate:"required"`
}

func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}