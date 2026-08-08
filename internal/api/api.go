package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"Geoapi/internal/logger"
	"Geoapi/internal/models"
)

const (
	apiBaseURL  = "https://catalog.api.2gis.com/3.0/items"
	staticMapURL = "https://static.maps.2gis.com/1.0"
	imageSize = "600x400"
	zoom = "17"
)

func GenerateStaticMapURL(lat, lon float64, apiKey string) (string) {
    if apiKey == "" {
        return "Can't get API!"
    }
    return fmt.Sprintf("%s?zoom=%s&size=%s&center=%f,%f&key=%s",
        staticMapURL, zoom, imageSize, lon, lat, apiKey)
}

func SearchLocations(query string, apiKey string) ([]models.Location, error) { 
	logger := logger.Logger()
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
	var apiResponse models.APIResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	logger.Info().Msgf("Parsed locations %d: ", len(apiResponse.Result.Items))
	return apiResponse.Result.Items, nil
}
