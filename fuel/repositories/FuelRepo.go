package repositories

import (
	Config "aiforesee/config"
	"aiforesee/models"

	_ "github.com/go-sql-driver/mysql"
)

func Fetch(fuels *[]models.Fuel) (err error) {

	if err = Config.DB.Debug().Order("fuel_liter ASC").Find(fuels).Error; err != nil {
		return err
	}
	return nil
}

func Detail(fuel *models.Fuel, id int32) (err error) {

	if err = Config.DB.Where("id = ?", id).First(fuel).Error; err != nil {
		return err
	}
	return nil
}

func Create(fuel *models.Fuel) (err error) {

	if err = Config.DB.Create(fuel).Error; err != nil {
		return err
	}
	return nil
}

func Update(fuel *models.Fuel) (err error) {
	if err = Config.DB.Model(&fuel).Updates(fuel).Error; err != nil {
		return err
	}
	return nil
}

func Delete(fuel *models.Fuel) (err error) {
	if err = Config.DB.Where("id = ?", fuel.ID).First(fuel).Error; err != nil {
		return err
	}
	if err = Config.DB.Delete(fuel).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAll() (err error) {
	var fuel models.Fuel
	if err = Config.DB.Delete(fuel).Error; err != nil {
		return err
	}
	return nil
}
