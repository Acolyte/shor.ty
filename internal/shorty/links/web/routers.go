package web

import (
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/access"
	"github.com/go-ozzo/ozzo-routing/v2/fault"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/go-ozzo/ozzo-routing/v2/slash"
	"log"
	"net/http"
	"shorty/internal/shorty/config"
)

func Serve() {
	router := routing.New()

	router.Get("/css/*", file.Server(file.PathMap{
		"/": "/web/static/",
	}))
	router.Get("/js/*", file.Server(file.PathMap{
		"/": "/web/static/",
	}))
	router.Get("/", IndexHandler)
	router.Post("/", CreateHandler)

	router.Get("/<id:[0-9a-zA-Z]*>", LinkByUUIDHandler)

	router.Use(
		access.Logger(log.Printf),
		slash.Remover(http.StatusMovedPermanently),
		fault.Recovery(log.Printf),
	)

	http.Handle("/", router)
	log.Printf("Server is running at %s:%d", config.Settings.Server.Host, config.Settings.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Settings.Server.Host, config.Settings.Server.Port), nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
