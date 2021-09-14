package myProject

import (
	"encoding/json"
)

type CityKv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MyCitiesModel struct {
	ID   uint64
	Name CityKvField `json:"my_city_name"`
}

func (city *CityKvField) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &city)
}
