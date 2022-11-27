package controllers

import (
	"errors"
	"fmt"
	"gin-api-template/models"
	"gin-api-template/responses"

	"github.com/gin-gonic/gin"
	"github.com/marcantoineg/windchill/windchill"
	"github.com/marcantoineg/windchill/windchill/speed"
	temp "github.com/marcantoineg/windchill/windchill/temperature"
)

func WindChill(c *gin.Context) {
	body, err := GetAndValidateBody[models.WindchillRequest](c)
	if err != nil {
		return
	}

	if err = validateBody(c, *body); err != nil {
		return
	}

	windchillIndex := windchill.GetWindChillIndex(body.Temperature, body.WindSpeed)
	SendJSON(c, responses.OK(windchillIndex))
}

func validateBody(c *gin.Context, b models.WindchillRequest) error {
	var err error = nil

	if !validateTemperatureUnit(b.Temperature) {
		err = errors.New(fmt.Sprintf("could not parse Temperature Unit: '%s'", b.Temperature.Unit))
		SendJSON(c, responses.BadRequest(err))
	} else if !validateSpeedUnit(b.WindSpeed) {
		err = errors.New(fmt.Sprintf("could not parse Speed Unit: '%s'", b.WindSpeed.Unit))
		SendJSON(c, responses.BadRequest(err))
	}

	return err
}

func validateTemperatureUnit(t temp.Temperature) bool {
	switch t.Unit {
	case temp.Celsius, temp.Fahrenheit, temp.Kelvin:
		return true
	default:
		return false
	}
}

func validateSpeedUnit(s speed.Speed) bool {
	switch s.Unit {
	case speed.KPH, speed.MPH:
		return true
	default:
		return false
	}
}
