package main

import (
	"fmt"
	"math/rand"
)

func main() {

	restaurants := []string{"Denny's @ Palomar", "Denny's @ Zapopan"}
	moreRestaurants := []string{"Panda Express", "Carl's JR", "Dave's Hot Chicken"}

	restaurants = append(restaurants, moreRestaurants...)

	randomIndex := rand.Intn(len(restaurants))

	fmt.Printf("----- You should eat at: %s -----\n", restaurants[randomIndex])

}
