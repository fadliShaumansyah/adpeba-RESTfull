package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv(){
    err := godotenv.Load()
    if err !=nil{
        panic("gagal load file .env")
    }
}

func GetDSN() string{
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
user, pass, host, port, name)
}


// ConnectDatabase menerima DSN dan mengembalikan *gorm.DB
func ConnectDatabase(dsn string) (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//  set global var
	DB = database

	return database, nil
}
