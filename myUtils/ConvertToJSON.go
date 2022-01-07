package myProject

import (
	"encoding/json"
)

type City struct {
	Id         int
	Name       string
	Population int
}

type CityKv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MyCitiesModel struct {
	ID         int64  `json:"id"`
	Name       string `json:"my_city_name"`
	Population string `json:"city_population"`
}

func (city *CityKv) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &city)
}
