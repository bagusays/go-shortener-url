package shortener

import (
	"database/sql"
	"fmt"
	"html"
	"url-shortener/models"

	"github.com/jmoiron/sqlx"
)

type ShortenerRepo struct {
	db *sqlx.DB
}

func NewShortenerRepo(db *sqlx.DB) *ShortenerRepo {
	return &ShortenerRepo{db: db}
}

func (r *ShortenerRepo) FindByShortLink(shortLink string) (*models.Link, error) {
	s := "SELECT * FROM links WHERE BINARY short_url = ?"
	var link models.Link
	err := r.db.Get(&link, s, shortLink)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(fmt.Errorf("shortener.repo.FindByShortLink: %s", err.Error()))
		return &link, err
	}

	return &link, nil
}

func (r *ShortenerRepo) CreateLink(data models.Link) error {
	q := "INSERT INTO links (short_url, long_url) VALUES (?, ?)"
	_, err := r.db.Exec(q, html.EscapeString(data.ShortUrl), html.EscapeString(data.LongUrl))
	if err != nil {
		fmt.Println(fmt.Errorf("shortener.repo.CreateLink: %s", err.Error()))
		return err
	}

	return nil
}

func (r *ShortenerRepo) EditLink(data models.Link) error {
	q := "UPDATE links SET long_url = ? WHERE short_url = ?"
	_, err := r.db.Exec(q, html.EscapeString(data.LongUrl), html.EscapeString(data.ShortUrl))
	if err != nil {
		fmt.Println(fmt.Errorf("shortener.repo.EditLink: %s", err.Error()))
		return err
	}

	return nil
}

func (r *ShortenerRepo) DeleteLink(data models.DeleteLink) error {
	q := "DELETE FROM links WHERE short_url = ?"
	_, err := r.db.Exec(q, html.EscapeString(data.ShortUrl))
	if err != nil {
		fmt.Println(fmt.Errorf("shortener.repo.DeleteLink: %s", err.Error()))
		return err
	}

	return nil
}
