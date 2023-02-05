package main

import (
	"fmt"

	controllers "example.com/go-crud-api/controller"
	"example.com/go-crud-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/meds/:id", controllers.ReadMed)
	r.GET("/meds", controllers.ReadMeds)
	r.POST("/meds", controllers.CreateMed)
	r.PUT("/meds/:id", controllers.UpdateMed)
	r.DELETE("/meds/:id", controllers.DeleteMed)
	r.Run(":5000")
}
