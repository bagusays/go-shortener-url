package shortener

import (
	"testing"
	"url-shortener/mock"
	"url-shortener/models"

	"github.com/stretchr/testify/assert"
)

func TestServiceFindByShortLink(t *testing.T) {
	m := new(mock.MockShortenerRepo)
	productExpected := models.Link{
		ShortUrl: "sqlx",
		LongUrl:  "https://godoc.org/github.com/julienschmidt/httprouter",
	}

	m.On("FindByShortLink").Return(&productExpected, nil)

	s := NewShortenerService(m)
	products, err := s.FindByShortLink(productExpected.ShortUrl)

	m.AssertExpectations(t)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, productExpected, *products)
}

func TestServiceCreateLink(t *testing.T) {
	m := new(mock.MockShortenerRepo)
	data := models.Link{
		LongUrl:  "http://jmoiron.github.io/sqlx/",
		ShortUrl: "sqlx",
	}

	m.On("CreateLink").Return(nil)

	s := NewShortenerService(m)
	err := s.CreateLink(data)

	m.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestServiceEditLink(t *testing.T) {
	m := new(mock.MockShortenerRepo)
	data := models.Link{
		LongUrl:  "http://jmoiron.github.io/sqlx/",
		ShortUrl: "sqlx",
	}

	m.On("EditLink").Return(nil)

	s := NewShortenerService(m)
	err := s.EditLink(data)

	m.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestServiceDeleteLink(t *testing.T) {
	m := new(mock.MockShortenerRepo)
	data := models.DeleteLink{
		ShortUrl: "sqlx",
	}

	m.On("DeleteLink").Return(nil)

	s := NewShortenerService(m)
	err := s.DeleteLink(data)

	m.AssertExpectations(t)

	assert.Nil(t, err)
}
