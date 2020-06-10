package main

import (
	"url-shortener/app/shortener"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func Router(r *httprouter.Router, db *sqlx.DB) {

	handlerShortener := shortener.NewHandlerShortener(db)

	r.GET("/:shortLink", handlerShortener.FindByShortLink())
	r.POST("/link/create", handlerShortener.CreateLink())
	r.POST("/link/edit", handlerShortener.EditLink())
	r.POST("/link/delete", handlerShortener.DeleteLink())
}
