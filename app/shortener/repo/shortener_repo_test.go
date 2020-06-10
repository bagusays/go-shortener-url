package shortener

import (
	"testing"
	"url-shortener/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepoFindByShortLink(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	shortUrl := "sqlx"
	rows1 := sqlmock.NewRows([]string{"short_url", "long_url"}).
		AddRow("sqlx", "https://godoc.org/github.com/julienschmidt/httprouter")

	query1 := "^SELECT (.+) FROM links WHERE BINARY short_url = \\?"

	mock.ExpectQuery(query1).WithArgs(shortUrl).WillReturnRows(rows1)

	productsExpected := models.Link{
		ShortUrl: shortUrl,
		LongUrl:  "https://godoc.org/github.com/julienschmidt/httprouter",
	}

	r := NewShortenerRepo(sqlxDB)

	productOK, err := r.FindByShortLink(productsExpected.ShortUrl)
	assert.NoError(t, err)
	assert.Equal(t, productsExpected, *productOK)
}

func TestRepoFindByShortLink_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	shortUrl := "sqlx"
	rows1 := sqlmock.NewRows([]string{"short_url", "long_url"})

	query1 := "^SELECT (.+) FROM links WHERE BINARY short_url = \\?"

	mock.ExpectQuery(query1).WithArgs(shortUrl).WillReturnRows(rows1)

	var productsExpected *models.Link = nil

	r := NewShortenerRepo(sqlxDB)

	productsOK, err := r.FindByShortLink(shortUrl)
	assert.NoError(t, err)
	assert.Equal(t, productsExpected, productsOK)
}

func TestRepoCreateLink(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	args := models.Link{
		LongUrl:  "https://stackoverflow.com/questions/57729947/how-to-test-mysql-insert-method",
		ShortUrl: "sqlx",
	}

	_ = mock.ExpectExec("^INSERT INTO links*").WithArgs(args.ShortUrl, args.LongUrl).WillReturnResult(sqlmock.NewResult(1, 1))

	r := NewShortenerRepo(sqlxDB)
	err = r.CreateLink(args)

	assert.Nil(t, err)
}

func TestRepoEditLink(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	args := models.Link{
		LongUrl:  "https://stackoverflow.com/questions/57729947/how-to-test-mysql-insert-method",
		ShortUrl: "sqlx",
	}

	_ = mock.ExpectExec("^UPDATE links SET long_url*").WithArgs(args.LongUrl, args.ShortUrl).WillReturnResult(sqlmock.NewResult(0, 1))

	r := NewShortenerRepo(sqlxDB)
	err = r.EditLink(args)

	assert.Nil(t, err)
}

func TestRepoDeleteLink(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	args := models.DeleteLink{
		ShortUrl: "sqlx",
	}

	_ = mock.ExpectExec("^DELETE FROM links*").WithArgs(args.ShortUrl).WillReturnResult(sqlmock.NewResult(0, 1))

	r := NewShortenerRepo(sqlxDB)
	err = r.DeleteLink(args)

	assert.Nil(t, err)
}
