package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Circle struct {
	Center LatLng  `json:"center"`
	Radius float64 `json:"radius"`
}

type LocationRestriction struct {
	Circle Circle `json:"circle"`
}

type PlacesAPIRequest struct {
	IncludedTypes       []string            `json:"includedTypes"`
	MaxResultCount      int                 `json:"maxResultCount"`
	LocationRestriction LocationRestriction `json:"locationRestriction"`
}

type Restaurants struct {
}

func GetRestaurant(coordinates Coordinates, url string, apiKey string, params string) ([]byte, error) {
	reqBody := PlacesAPIRequest{
		IncludedTypes:  []string{"restaurant"},
		MaxResultCount: 10,
		LocationRestriction: LocationRestriction{
			Circle: Circle{
				Center: LatLng{
					Latitude:  coordinates.Latitude,
					Longitude: coordinates.Longitude,
				},
				Radius: 500.0,
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", apiKey)
	req.Header.Set("X-Goog-FieldMask", params)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
