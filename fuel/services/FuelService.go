package services

import (
	"aiforesee/fuel/repositories"
	"aiforesee/models"
)

func ListFuel() (dataRet []models.Fuel, err error) {
	err = repositories.Fetch(&dataRet)
	return dataRet, err
}

func DetailFuel(id int32) (dataRet models.Fuel, err error) {
	err = repositories.Detail(&dataRet, id)
	return dataRet, err
}

func CreateFuel(fuel models.Fuel) (dataRet models.Fuel, err error) {
	err = repositories.Create(&fuel)
	return fuel, err
}

func UpdateFuel(fuel models.Fuel) (dataRet models.Fuel, err error) {
	err = repositories.Update(&fuel)
	return fuel, err
}

func DeleteFuel(fuel models.Fuel) (dataRet models.Fuel, err error) {
	err = repositories.Delete(&fuel)
	return fuel, err
}
func GenerateDataFuel() {
	repositories.DeleteAll()
	basePricePremium := 6450
	basePricePertalite := 7650
	liter := 1
	for liter <= 20 {
		fuelPremium := models.Fuel{
			FuelType:  "Premium",
			FuelPrice: liter * basePricePremium,
			FuelLiter: liter,
		}
		fuelPertalite := models.Fuel{
			FuelType:  "Pertalite",
			FuelPrice: liter * basePricePertalite,
			FuelLiter: liter,
		}
		repositories.Create(&fuelPremium)
		repositories.Create(&fuelPertalite)
		liter++
	}
}
