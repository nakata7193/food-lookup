package controllers

import (
	"foodLookup/cmd/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFood(data model.FoodRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		foodID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "conversion error"})
		}
		food, err := data.GetFoodByID(foodID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}
		ctx.JSON(http.StatusOK, food)
	}
}

func GetAllFoods(data model.FoodRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		foods, err := data.GetAllFoods()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "get all foods"})
		}
		ctx.JSON(http.StatusOK, foods)
	}
}

func AddFood(data model.FoodRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		food := model.Food{}
		ctx.BindJSON(&food)
		err := data.AddFood(&food)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid food"})
		}
		ctx.JSON(http.StatusAccepted, food)
	}
}
