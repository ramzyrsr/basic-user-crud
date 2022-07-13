package database

import "fmt"

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}

func InitDB() {
	config := Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "password",
		DB:         "lemonilo",
	}

	connectString := GetConnectionString(config)
	err := Connect(connectString)
	if err != nil {
		panic(err.Error())
	}
}
