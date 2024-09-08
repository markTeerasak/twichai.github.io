package main

import "github.com/twichai/twichai.github.io/database"

func main() {
	db := database.ConnectDB()
	database.Migrate(db)
}
