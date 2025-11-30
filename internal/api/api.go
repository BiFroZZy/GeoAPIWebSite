package api

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"encoding/json"
	"net/http"
	"io"
)

const (
	apiBaseURL  = "https://catalog.api.2gis.com/3.0/items"
	staticMapURL = "https://static.maps.2gis.com/1.0"
	imageSize = "600x400"
	zoom = "17"
)

type Point struct{
	Lat 	 float64  `json:"lat"`	// создали структуру для координат потом засунули в Location 
	Lon 	 float64  `json:"lon"`
}

type Location struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address_name"`
	Point 	 Point	  `json:"point"` // все координаты
}

type APIResponse struct {
	Result struct {
		Items []Location `json:"items"`
	} `json:"result"`
}

func GenerateStaticMapURL(lat, lon float64) (string) {
    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        return "Can't get API!"
    }
    return fmt.Sprintf("%s?zoom=%s&size=%s&center=%f,%f&key=%s",
        staticMapURL, zoom, imageSize, lon, lat, apiKey)
}

func SearchLocations(query string) ([]Location, error) { // добавил lon lat в функцию
	if err := godotenv.Load(); err != nil{log.Printf("Can't get .env file: %v", err)}
	 // получение API 
	apiKey := os.Getenv("API")

	url := fmt.Sprintf("%s?q=%s&key=%s&fields=items.point,items.address_name,items.photo_ids", apiBaseURL, query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call 2GIS API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("2GIS API returned status %d: %s", resp.StatusCode, body)
	}
	var apiResponse APIResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	log.Printf("Parsed %d locations", len(apiResponse.Result.Items))
	return apiResponse.Result.Items, nil
}
