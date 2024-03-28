package main

import (
	"foodLookup/cmd"
	"foodLookup/cmd/controllers"
	"foodLookup/cmd/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := db.DbInit()
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	err := repository.CreateTableIfNotExists()
	if err != nil {
		log.Fatalf("Error creating 'foods' table: %v", err)
	}

	router.GET("/health", controllers.HealthHandler())

	router.GET("/api/food/:id", controllers.GetFood(*repository))

	router.GET("api/foods", controllers.GetAllFoods(*repository))

	router.POST("api/food", controllers.AddFood(*repository))

	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
