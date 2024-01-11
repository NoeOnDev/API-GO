package models

import "github.com/go-pg/pg/v10/orm"

type User struct {
    tableName struct{} `pg:"users"` 
    ID        int      `pg:"id,pk"` 
    FirstName string   `pg:"first_name"`
    LastName  string   `pg:"last_name"`
    Username  string   `pg:"username,unique"` 
    Phone     string   `pg:"phone"`
    Email     string   `pg:"email,unique"` 
    Password  string   `pg:"password"`    
}
