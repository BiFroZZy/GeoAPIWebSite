package server

import (
	"context"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Geoapi/internal/configs"
	"Geoapi/web/pagehandlers"
	logs "Geoapi/internal/logger"
)

var logger = logs.Logger()

type Server struct{
	ctx context.Context
	e *echo.Echo
	cfg *configs.Configs
}

type Template struct{
	templates *template.Template // Структура для шаблонов
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	return t.templates.ExecuteTemplate(w, name, data) // Метод для рендера шаблонов 
}

func New() (*Server, error){
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	ctx := context.Background()
	cfg, err := configs.Load()
	if err != nil{
		logger.Error().Err(err).Msg("Failed to get configs while starting server")
		return nil, err
	}
	return &Server{
		ctx: ctx,
		e: e,
		cfg: cfg,
	}, nil
}

func (s *Server) SetMiddleware(){
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
}

func (s *Server) SetRoutes(){
	s.e.GET("/", pagehandlers.HomePage)
	s.e.GET("/settings", pagehandlers.SettingsPage)
	s.e.GET("/about", pagehandlers.AboutPage)
	s.e.GET("/help", pagehandlers.HelpPage)
	s.e.GET("/info", pagehandlers.HomePage)
}

func (s *Server) SetTemplates(){
	templates, err := template.ParseFiles( 
		"web/templates/footer.html",
	    "web/templates/header.html",
		"web/templates/about_page.html",
		"web/templates/home_page.html",
		"web/templates/help_page.html",
		"web/templates/search_page.html",
		"web/templates/settings_page.html",
	); 
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get templates")
	}

	s.e.Renderer = &Template{templates: templates}

	fs := http.FileServer(http.Dir("static")) // хранение статитечских данных
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	s.e.Static("/web/static", "web/static") 
}
func (s *Server) Start() error{
	s.SetMiddleware()
	s.SetTemplates()
	s.SetRoutes()
	if err := s.e.Start(s.cfg.ServerPort); err !=nil{
		logger.Error().Err(err).Msg("Failed to start server")
		return err
	}
	return nil
}