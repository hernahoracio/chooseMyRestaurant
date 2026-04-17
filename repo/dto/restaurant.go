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

type CreateRestaurantInput struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Location  string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Cuisine   string  `json:"cuisine"`
	Rating    float64 `json:"rating"`
}

type RestaurantQuery struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Location  string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Cuisine   string  `json:"cuisine"`
	Rating    float64 `json:"rating"`
}

// API to DTO format
func (r *RestaurantQuery) FromModel(m models.CreateRestaurantInput) {
	r.ID = m.ID
	r.Name = m.Name
	r.Location = m.Location
	r.Latitude = m.Latitude
	r.Longitude = m.Longitude
	r.Cuisine = m.Cuisine
	r.Rating = m.Rating

}

func (r *CreateRestaurantInput) FromModel(m models.CreateRestaurantInput) {
	r.ID = m.ID
	r.Name = m.Name
	r.Location = m.Location
	r.Latitude = m.Latitude
	r.Longitude = m.Longitude
	r.Cuisine = m.Cuisine
	r.Rating = m.Rating

}

// Dto to API response
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

type CreateRestaurantQuery struct {
	ID        string
	Name      string
	Location  string
	Latitude  float64
	Longitude float64
	Cuisine   string
	Rating    float64
}

type FindRestaurantQuery struct {
	Name     string
	Location string
	Cuisine  string
	Rating   float64
}
