package main

import (
	"fmt"
	"log"

	"github.com/your-username/your-project/petstore"
	"github.com/your-username/your-project/petstore/api"
)

func main() {
	cfg := petstore.NewConfiguration()
	client := petstore.NewAPIClient(cfg)

	// List pets
	pets, _, err := client.PetsApi.GetPets(nil)
	if err != nil {
		log.Fatalf("Error getting pets: %v", err)
	}
	for _, pet := range pets {
		fmt.Printf("Pet ID: %d, Name: %s\n", pet.Id, pet.Name)
	}

	// Create a new pet
	newPet := api.Pet{
		Id:   3,
		Name: "Birdy",
	}
	createdPet, _, err := client.PetsApi.CreatePet(nil, newPet)
	if err != nil {
		log.Fatalf("Error creating pet: %v", err)
	}
	fmt.Printf("Created pet: ID %d, Name %s\n", createdPet.Id, createdPet.Name)
}
