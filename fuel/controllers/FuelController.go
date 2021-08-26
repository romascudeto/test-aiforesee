package controllers

import (
	"aiforesee/fuel/services"
	"aiforesee/helper"
	"aiforesee/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GetFuel ... Get all fuel
func GetListFuel(c echo.Context) error {

	data, err := services.ListFuel()
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//GetFuelByID ... Get all fuel by id
func GetFuelByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := services.DetailFuel(int32(idInt))
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusNotFound, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//CreateFuel ... Create fuel
func CreateFuel(c echo.Context) error {
	var dataFuel models.Fuel
	c.Bind(&dataFuel)
	createdData, err := services.CreateFuel(dataFuel)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, createdData)
		return c.JSON(http.StatusOK, resp)
	}
}

//UpdateFuel ... Update fuel
func UpdateFuel(c echo.Context) error {
	var dataFuel models.Fuel
	c.Bind(&dataFuel)
	updatedData, err := services.UpdateFuel(dataFuel)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, updatedData)
		return c.JSON(http.StatusOK, resp)
	}
}

//DeleteFuel ... Delete fuel
func DeleteFuel(c echo.Context) error {
	var dataFuel models.Fuel
	c.Bind(&dataFuel)
	deletedData, err := services.DeleteFuel(dataFuel)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, deletedData)
		return c.JSON(http.StatusOK, resp)
	}
}
