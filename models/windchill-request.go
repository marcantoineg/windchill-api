package models

import (
	"github.com/marcantoineg/windchill/windchill/speed"
	temp "github.com/marcantoineg/windchill/windchill/temperature"
)

type WindchillRequest struct {
	Temperature temp.Temperature `json:"temperature"`
	WindSpeed   speed.Speed      `json:"windSpeed"`
}
