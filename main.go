package main

import (
	"fmt"
	"math/rand"
)

func main() {

	restaurants := []string{"Denny's @ Palomar", "Denny's @ Zapopan"}

	randomIndex := rand.Intn(len(restaurants))

	fmt.Printf("-----You should eat at this Denny's: %s -----\n", restaurants[randomIndex])

}
