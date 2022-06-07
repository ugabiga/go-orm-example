package config

import "fmt"

func GetDBInfo() (string, string) {
	return "postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "user", "pass", "postgres")
}
