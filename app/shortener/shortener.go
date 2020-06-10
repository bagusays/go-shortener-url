package shortener

import (
	controller "url-shortener/app/shortener/controllers"
	repo "url-shortener/app/shortener/repo"
	service "url-shortener/app/shortener/services"

	"github.com/jmoiron/sqlx"
)

type HandlerShortener struct {
	controller.ShortenerController
}

func NewHandlerShortener(db *sqlx.DB) HandlerShortener {
	repo := repo.NewShortenerRepo(db)
	service := service.NewShortenerService(repo)
	controller := controller.NewShortenerController(service)

	return HandlerShortener{controller}
}
