package model

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type FoodRepository interface {
	GetFoodByID(id int) (*Food, error)
	GetAllFoods() ([]*Food, error)
	AddFood(food *Food) error
}

type FoodRepo struct {
	db *sql.DB
}

func NewFoodRepo(db *sql.DB) *FoodRepo {
	return &FoodRepo{db: db}
}

func (r *FoodRepo) CreateTableIfNotExists() error {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS foods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR(255) NOT NULL,
    carbs DECIMAL(10,2) NOT NULL,
    protein DECIMAL(10,2) NOT NULL,
    fat DECIMAL(10,2) NOT NULL
)`)
	if err != nil {
		return err
	}
	return nil
}

func (r *FoodRepo) GetFoodByID(id int) (*Food, error) {
	food := &Food{}
	err := r.db.QueryRow("SELECT * FROM foods WHERE id = ?", id).Scan(&food.ID, &food.Description, &food.Carbs, &food.Protein, &food.Fat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return food, nil
}

func (r *FoodRepo) GetAllFoods() ([]*Food, error) {
	rows, err := r.db.Query("SELECT * FROM foods")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []*Food
	for rows.Next() {
		food := &Food{}
		err := rows.Scan(&food.ID, &food.Description, &food.Carbs, &food.Protein, &food.Fat)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foods, nil
}

func (r *FoodRepo) AddFood(food *Food) error {
	_, err := r.db.Exec("INSERT INTO foods (description, carbs, protein, fat) VALUES (?, ?, ?, ?)", food.Description, food.Carbs, food.Protein, food.Fat)
	if err != nil {
		return err
	}
	return nil
}
