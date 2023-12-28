package main

import (
	"net/http"

	"github.com/servusdei2018/shawty/pkg/conf"
	"github.com/servusdei2018/shawty/pkg/db"
	"github.com/servusdei2018/shawty/pkg/routes"
)

func main() {
	config := conf.Parse()

	var database db.DB
	if err := database.Init(); err != nil {
		panic(err)
	}

	r := routes.New(&database, config.Auth)

	if config.Port == "" {
		config.Port = "8080"
	}
	if err := http.ListenAndServe("0.0.0.0:"+config.Port, r.Mux); err != nil {
		panic(err)
	}
}
