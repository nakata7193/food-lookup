package model

type Food struct {
	Id	int json:"id"
	Description	string json:"description"
	Kcal	float64 		json:"kcal"
	Protein	float64
	Fat	float64
	Carbs	float64
}