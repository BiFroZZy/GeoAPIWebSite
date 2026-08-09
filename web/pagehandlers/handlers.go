package pagehandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"Geoapi/internal/api"
	"Geoapi/internal/configs"
	"Geoapi/internal/models"
)

func AboutPage(c echo.Context) error{
	return c.Render(http.StatusOK, "about_page", map[string]interface{}{
		"Title": "About",
	})
}

func HomePage(c echo.Context) error{
	return c.Render(http.StatusOK, "home_page", map[string]interface{}{
		"Title": "Home",
	})
}

func SearchPage(c echo.Context) error{
	cfg, err := configs.Load()
	if err != nil{
		return err
	}
	query := c.FormValue("q")
	if query == "" {
		return c.Render(http.StatusBadRequest, "search_page", map[string]interface{}{
			"Title": "Search",
			"Error": "Введите поисковый запрос",
			"Query": "Nothing is here",
		})
	}

	locations, err := api.SearchLocations(query, cfg.APIkey) 
	if err != nil {
		c.Logger().Errorf("Search API error: %v", err)
		return c.Render(http.StatusInternalServerError, "search_page", map[string]interface{}{
			"Title": "Search",
			"Error": "Произошла ошибка при поиске. Попробуйте позже.",
			"Query": query,
		})
	}
	
	var viewLocations []models.ViewLocation
	for _, loc := range locations{ 
		mapURL := api.GenerateStaticMapURL(loc.Point.Lat, loc.Point.Lon, cfg.APIkey) // функция которая генерирует URL для статической карты
		vl := models.ViewLocation{
			ID: 		loc.ID,
			Name: 		loc.Name,
			Address: 	loc.Address,
			Lat: 		loc.Point.Lat,
			Lon: 		loc.Point.Lon,
			MapURL:     mapURL,
		}
		
		err := models.Validation(vl)
		if err != nil{
			return err
		}
		viewLocations = append(viewLocations, vl)
	}

	data := struct {
		Query     string
		Locations []models.ViewLocation
	}{
		Query:     query,
		Locations: viewLocations,
	}
	return c.Render(http.StatusOK, "search_page", data)
}