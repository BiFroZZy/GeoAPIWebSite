package pagehandlers

import (
	"net/http"
	"Geoapi/internal/api"
	
	"github.com/labstack/echo/v4"
)

func HelpPage(c echo.Context) error{
	return c.Render(http.StatusOK, "help_page", map[string]interface{}{
		"Title": "Help",
	})
}

func SettingsPage(c echo.Context) error{
	return c.Render(http.StatusOK, "settings_page", map[string]interface{}{
		"Title": "Settings",
	})
}

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
	// инициализация страницы
	query := c.FormValue("q")
	if query == "" {
		return c.Render(http.StatusBadRequest, "search_page", map[string]interface{}{
			"Title": "Search",
			"Error": "Введите поисковый запрос",
			"Query": "",
		})
	}

	locations, err := api.SearchLocations(query) 
	if err != nil {
		c.Logger().Errorf("Search API error: %v", err)
		return c.Render(http.StatusInternalServerError, "search_page", map[string]interface{}{
			"Title": "Search",
			"Error": "Произошла ошибка при поиске. Попробуйте позже.",
			"Query": query,
		})
	}
	type ViewLocation struct{
		ID 		string 		
		Name 	string 	
		Address string	
		Lat		float64 // обьявление координат
		Lon 	float64
		MapURL  string // статитческий URL карты
	}
	var viewLocations []ViewLocation
	for _, loc := range locations{ 
		mapURL := api.GenerateStaticMapURL(loc.Point.Lat, loc.Point.Lon) // функция которая генерирует URL для статической карты
		vl := ViewLocation{
			ID: 		loc.ID,
			Name: 		loc.Name,
			Address: 	loc.Address,
			Lat: 		loc.Point.Lat,
			Lon: 		loc.Point.Lon,
			MapURL:     mapURL,
		}
		viewLocations = append(viewLocations, vl)
	}

	data := struct {
		Query     string
		Locations []ViewLocation
	}{
		Query:     query,
		Locations: viewLocations,
	}
	return c.Render(http.StatusOK, "search_page", data)
}