package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Geoapi/web/pagehandlers"
)

type Template struct{
	templates *template.Template // Структура для шаблонов
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	return t.templates.ExecuteTemplate(w, name, data) // Метод для рендера шаблонов 
}

func HandleRequests(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	templates, err := template.ParseFiles( 
		"web/templates/footer.html",
	    "web/templates/header.html",
		"web/templates/about_page.html",
		"web/templates/home_page.html",
		"web/templates/help_page.html",
		"web/templates/search_page.html",
		"web/templates/settings_page.html",
	); 
	if err != nil {log.Fatalf("Ошибка загрузки шаблонов:%v", err)}

	if err != nil{
		log.Printf("Can't get the html-templates: %v", err)
	}
	e.Renderer = &Template{templates: templates}

	fs := http.FileServer(http.Dir("static")) // хранение статитечских данных
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	e.Static("/web/static", "web/static") 

	e.GET("/", pagehandlers.HomePage)
	e.GET("/settings_page", pagehandlers.SettingsPage)
	e.GET("/about_page", pagehandlers.AboutPage)
	e.GET("/help_page", pagehandlers.HelpPage)
	e.GET("/search_page", pagehandlers.HomePage)

	e.Logger.Fatal(e.Start(":8081"))
}
