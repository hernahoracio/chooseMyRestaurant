package dto

import "chooseMyRestaurant/models"

type Restaurant struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Location  string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Cuisine   string  `json:"cuisine"`
	Rating    float64 `json:"rating"`
}

type createRestaurantInput struct {
	ID        *string `json:"id"`
	Name      string  `json:"name"`
	Location  string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Cuisine   string  `json:"cuisine"`
	Rating    float64 `json:"rating"`
}

func (r *Restaurant) ToModel() models.Restaurant {
	m := models.Restaurant{
		ID:        r.ID,
		Name:      r.Name,
		Location:  r.Location,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		Cuisine:   r.Cuisine,
		Rating:    r.Rating,
	}
	return m
}

type createRestaurantQuery struct {
	ID        string
	Name      string
	Location  string
	Latitude  float64
	Longitude float64
	Cuisine   string
	Rating    string
}

type findRestaurantQuery struct {
	Name     string
	Location string
	Cuisine  string
	Rating   float64
}
