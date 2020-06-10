package domain

import "url-shortener/models"

type ShortenerRepoInterface interface {
	FindByShortLink(string) (*models.Link, error)
	EditLink(models.Link) error
	CreateLink(models.Link) error
	DeleteLink(models.DeleteLink) error
}

type ShortenerServiceInterface interface {
	FindByShortLink(string) (*models.Link, error)
	EditLink(models.Link) error
	CreateLink(models.Link) error
	DeleteLink(models.DeleteLink) error
}
