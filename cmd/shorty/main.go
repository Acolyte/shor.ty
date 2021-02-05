package main

import (
	"shorty/internal/app/shorty"
	"shorty/internal/app/shorty/api"
)

func main() {

	err := shorty.Migrate()
	if err != nil {
		panic(err)
	}

	api.Serve()
}
