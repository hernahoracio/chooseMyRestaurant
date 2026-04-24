package models

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

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

func (rQ *RestaurantQuery) Query() (url.Values, error) {
	return query.Values(rQ)
}

func (r Restaurant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Restaurant) UnmarshallJSON(data []byte) (string, error) {
	var input string
	if err := json.Unmarshal(data, &input); err != nil {
		return "", err
	}
	return input, nil
}
