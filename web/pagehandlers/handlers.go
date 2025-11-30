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
		http.Error(c.Response(), "Нужно что-то ввести", http.StatusBadRequest)
		return c.Render(http.StatusOK, "search_page", map[string]interface{}{
		"Title": "Home",
		"Error": "Введите что-то",
	})
	}
	locations, err := api.SearchLocations(query) 
	if err != nil{
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
		return nil
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