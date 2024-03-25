package model

type Food struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Kcal        float64 `json:"kcal"`
	Protein     float64 `json:"protein"`
	Fat         float64 `json:"fat"`
	Carbs       float64 `json:"carbs"`
}
