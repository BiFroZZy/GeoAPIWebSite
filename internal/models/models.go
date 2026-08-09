package models

import (
	"github.com/go-playground/validator/v10"
	logs "Geoapi/internal/logger"
)

var (
	validate = validator.New()
	logger = logs.Logger()
)

type Point struct{
	Lat 	 	float64  `json:"lat"`	
	Lon 	 	float64  `json:"lon"`
}

type Location struct {
	ID       	string   `json:"id"`
	Name     	string   `json:"name"`
	Address  	string   `json:"address_name"`
	Point 	 	Point	 `json:"point"` 		// все координаты
}

type ViewLocation struct{
	ID 			string 		`validate:"required"`
	Name 		string 		`validate:"required,min=2,max=20"`
	Address 	string		`validate:"required,min=2,max=30"`
	Lat			float64 	`validate:"required"`	// обьявление координат
	Lon 		float64		`validate:"required"`
	MapURL  	string 		`validate:"required,url"`	// статитческий URL карты
}

type APIResponse struct {
	Result struct {
		Items []Location `json:"items"`
	} `json:"result"`
}

func Validation(s interface{}) error{
	err := validate.Struct(s)
	if err != nil{
		logger.Error().Err(err).Msg("Failed to validate struct")
		return err
	}
	return nil
}