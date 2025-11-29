package pagehandlers

import (
	"net/http"
	"html/template"
	"Geoapi/internal/api"
)

func HelpPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("web/templates/help_page.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "help_page", nil)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SettingsPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("web/templates/settings_page.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	err = tmpl.ExecuteTemplate(w, "settings_page", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) 
	}
}

func AboutPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("web/templates/about_page.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	err = tmpl.ExecuteTemplate(w, "about_page", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) 
	}
}

func HomePage(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r) // добавлено только что 
		return
	}
	
	tmpl, err := template.ParseFiles("web/templates/home_page.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	err = tmpl.ExecuteTemplate(w, "home_page", nil)
	if err != nil{ 
		http.Error(w, err.Error(), http.StatusInternalServerError) 
	}
}

func SearchPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("web/templates/search_page.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} // инициализация страницы

	query := r.URL.Query().Get("q") // считывает данные со строки поиска
	if query == "" {
		http.Error(w, "Нужно что-то ввести", http.StatusBadRequest)
		return
	}
	
	locations, err := api.SearchLocations(query) 
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	err = tmpl.ExecuteTemplate(w, "search_page", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}