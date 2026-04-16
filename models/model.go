package models

import "encoding/json"

type Restaurant struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Location  string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Cuisine   string  `json:"cuisine"`
	Rating    float64 `json:"rating"`
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
