package shortener

import (
	"fmt"
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
	testCase := []struct {
		mockFindByShortLink *models.Link
		shortUrl            string
		longUrl             string
		expected            interface{}
	}{
		{nil, "sqlx", "http://jmoiron.github.io/sqlx/", nil},
		{&models.Link{}, "", "", fmt.Errorf("SHORT_URL_IS_TAKEN")},
	}

	for _, data := range testCase {
		m := new(mock.MockShortenerRepo)
		link := models.Link{
			LongUrl:  data.longUrl,
			ShortUrl: data.shortUrl,
		}

		// var link *models.Link

		m.On("CreateLink").Return(nil)
		m.On("FindByShortLink").Return(data.mockFindByShortLink, nil)

		s := NewShortenerService(m)
		err := s.CreateLink(link)

		// m.AssertExpectations(t)

		// assert.Nil(t, err)
		assert.Equal(t, data.expected, err)
	}
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
