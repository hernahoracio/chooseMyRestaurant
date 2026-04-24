package main

import (
	"chooseMyRestaurant/internal"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	provider := internal.CoreLocationCLI{}
	coords, err := provider.GetCoordinates()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	body, err := internal.GetRestaurant(coords, cfg.PlacesBaseURL, cfg.API_Key, cfg.Parameters)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Restaurants near you: %s", body)

}
