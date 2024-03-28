package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type User struct {
	ID        string `json:"id" gorm:"primarykey"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func InitPostgresDB() {
	env := os.Getenv("ENV")

	if env != "prod" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(User{})
}

func CreateUser(user *User) (*User, error) {
	user.ID = uuid.New().String()
	res := db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func GetUser(id string) (*User, error) {
	var user User
	res := db.First(&user, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("user of id %s not found", id)
	}
	return &user, nil
}

func DeleteUser(id string) error {
	var deletedUser User
	result := db.Where("id = ?", id).Delete(&deletedUser)
	if result.RowsAffected == 0 {
		return errors.New("user not deleted")
	}
	return nil
}

func UpdateUser(user *User) (*User, error) {
	var userToUpdate User
	result := db.Model(&userToUpdate).Where("id = ?", user.ID).Updates(user)
	if result.RowsAffected == 0 {
		return &userToUpdate, errors.New("user not updated")
	}
	return user, nil
}
