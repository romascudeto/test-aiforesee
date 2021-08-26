package models

type Fuel struct {
	ID        int64
	FuelLiter int    `json:"fuel_liter"`
	FuelType  string `json:"fuel_type"`
	FuelPrice int    `json:"fuel_price"`
}

func (b *Fuel) TableName() string {
	return "fuels"
}
