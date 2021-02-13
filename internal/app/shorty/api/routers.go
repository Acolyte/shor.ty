package api

import (
	"fmt"
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/access"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/fault"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/go-ozzo/ozzo-routing/v2/slash"
	"log"
	"net/http"
	"shorty/internal/app/config"
	"shorty/internal/app/shorty"
)

// @title Shorty API
// @version 1.0
// @description API for working with shor.ty.
// @termsOfService http://swagger.io/terms/

// @contact.name API Development
// @contact.url https://github.com/Acolyte/shor.ty
// @contact.email acolytee@gmail.com

// @license.name Public

// @host shor.ty
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-AUTH-TOKEN

func Serve() {
	router := routing.New()

	router.Get("/swagger/swagger.json", file.Content("swagger/swagger.json"))
	router.Get("/swagger/swagger.yaml", file.Content("swagger/swagger.yaml"))
	router.Get("/css/*", file.Server(file.PathMap{
		"/": "/web/static/",
	}))
	router.Get("/js/*", file.Server(file.PathMap{
		"/": "/web/static/",
	}))
	router.Get("/", shorty.IndexHandler)
	router.Post("/", shorty.CreateHandler)

	router.Get("/<id:[0-9a-zA-Z]*>", shorty.LinkByUUIDHandler)

	apiV1 := router.Group(`/api/v1`)
	apiV1.Use(
		// all these handlers are shared by every route
		access.Logger(log.Printf),
		slash.Remover(http.StatusMovedPermanently),
		fault.Recovery(log.Printf),
		content.TypeNegotiator(content.JSON),
		HealthCheck,
		Encoding,
	)

	v1Links := apiV1.Group("/links")

	v1Links.Get(`/<id:\d+>`, shorty.LinkByIDHandler)
	v1Links.Get(``, shorty.LinksListHandler)
	v1Links.Post(``, shorty.LinkCreateHandler)
	v1Links.Delete(`/<id:\d+>`, shorty.LinkDeleteHandler)

	http.Handle("/", router)
	log.Printf("Server is running at %s:%d", config.Settings.Server.Host, config.Settings.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Settings.Server.Host, config.Settings.Server.Port), nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
