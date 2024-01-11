package models

import (
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID        int      `pg:"id,pk"`
    FirstName string   `pg:"firstname" validate:"required"`
    LastName  string   `pg:"lastname" validate:"required"`
    Username  string   `pg:"username,unique" validate:"required"`
    Phone     string   `pg:"phone" validate:"required"`
    Email     string   `pg:"email,unique" validate:"required,email"`
    Password  string   `pg:"password" validate:"required, min=6"`
}

func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}