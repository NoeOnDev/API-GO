package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"type:varchar(100);unique_index"`
}

var db *gorm.DB

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{})

    r := gin.Default()

    r.GET("/users", GetUsers)
    r.POST("/users", CreateUser)

    r.Run()
}

func GetUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error retrieving users"})
		return
	}
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": "Bad request"})
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(200, user)
}