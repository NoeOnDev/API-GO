package models

type User struct {
    TableName struct{} `pg:"users"`
    ID        int      `pg:"id,pk"`
    FirstName string   `pg:"first_name"`
    LastName  string   `pg:"last_name"`
    Username  string   `pg:"username,unique"`
    Phone     string   `pg:"phone"`
    Email     string   `pg:"email,unique"`
    Password  string   `pg:"password"`
}