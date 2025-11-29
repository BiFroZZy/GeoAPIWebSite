package handlers

import (
	"net/http"
	"Geoapi/web/pagehandlers"
)

func HandleRequests(){
	fs := http.FileServer(http.Dir("static")) // хранение статитечских данных
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	// обьявление страниц на хэдере
	http.HandleFunc("/", pagehandlers.HomePage)
	http.HandleFunc("/settings_page", pagehandlers.SettingsPage)
	http.HandleFunc("/about_page", pagehandlers.AboutPage)
	http.HandleFunc("/help_page", pagehandlers.HelpPage) 
	http.HandleFunc("/search_page", pagehandlers.SearchPage)

	http.ListenAndServe(":8081", nil)
}
