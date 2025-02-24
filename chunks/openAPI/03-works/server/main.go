package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email soberkoder@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/pets", getPets)
	r.POST("/pets", createPet)

	log.Fatal(r.Run())
}

type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// getPets godoc
// @Summary List pets
// @Description get pets
// @Tags pets
// @Accept  json
// @Produce  json
// @Success 200 {array} Pet
// @Router /pets [get]
func getPets(c *gin.Context) {
	pets := []Pet{
		{ID: 1, Name: "Doggy"},
		{ID: 2, Name: "Kitty"},
	}
	c.JSON(http.StatusOK, pets)
}

// createPet godoc
// @Summary Create pet
// @Description create pet
// @Tags pets
// @Accept  json
// @Produce  json
// @Param pet body Pet true "Create pet"
// @Success 200 {object} Pet
// @Router /pets [post]
func createPet(c *gin.Context) {
	var pet Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}
