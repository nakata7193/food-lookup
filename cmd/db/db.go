package db

import (
	"database/sql"
	"foodLookup/cmd/model"

	_ "modernc.org/sqlite"
)

func DbInit() *model.FoodRepo {
	db, err := sql.Open("sqlite", "food.db")
	if err != nil {
		panic(err)
	}
	repository := model.NewFoodRepo(db)
	return repository
}
