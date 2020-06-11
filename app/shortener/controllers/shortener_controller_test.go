package shortener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/mock"
	"url-shortener/models"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestControllerFindByShortLink(t *testing.T) {
	m := new(mock.MockShortenerService)
	productExpected := models.Link{
		ShortUrl: "sqlx",
		LongUrl:  "https://godoc.org/github.com/julienschmidt/httprouter",
	}

	m.On("FindByShortLink").Return(&productExpected, nil)

	router := httprouter.New()

	s := NewShortenerController(m)

	router.GET("/:shortLink", s.FindByShortLink())

	req, _ := http.NewRequest("GET", "/sqlx", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	m.AssertExpectations(t)

	assert.Equal(t, productExpected.LongUrl, rr.HeaderMap.Get("Location"))
}

func TestControllerFindByShortLink_NotFound(t *testing.T) {
	m := new(mock.MockShortenerService)
	var productExpected *models.Link = nil

	m.On("FindByShortLink").Return(productExpected, nil)

	router := httprouter.New()

	s := NewShortenerController(m)

	router.GET("/:name", s.FindByShortLink())

	req, _ := http.NewRequest("GET", "/wwwwwwwwwww", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	m.AssertExpectations(t)

	expected := `{"data":null,"message":"404 Not Found","success":false}`

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestControllerCreateLink(t *testing.T) {
	testCase := []struct {
		shortUrl string
		longUrl  string
		expected string
	}{
		{"json", "https://stackoverflow.com/questions/33659298/unit-testing-http-json-response-in-golang", fmt.Sprintf(`{"data":"http://localhost:8888/%s","message":null,"success":true}`, "json")},
		{"", "https://stackoverflow.com/questions/33659298/unit-testing-http-json-response-in-golang", `{"data":null,"message":"ShortURL must be provided","success":false}`},
		{"asd", "", `{"data":null,"message":"LongURL must be provided","success":false}`},
	}

	for _, data := range testCase {
		m := new(mock.MockShortenerService)
		m.On("CreateLink").Return(nil)

		router := httprouter.New()

		s := NewShortenerController(m)
		router.POST("/link/create", s.CreateLink())

		params := models.Link{
			LongUrl:  data.longUrl,
			ShortUrl: data.shortUrl,
		}

		j, err := json.Marshal(params)
		assert.Nil(t, err)

		req, _ := http.NewRequest("POST", "/link/create", bytes.NewReader(j))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, data.expected, rr.Body.String())
		// m.AssertExpectations(t)
	}

}

func TestControllerEditLink(t *testing.T) {
	testCase := []struct {
		shortUrl string
		longUrl  string
		expected string
	}{
		{"json", "https://stackoverflow.com/questions/33659298/unit-testing-http-json-response-in-golang", fmt.Sprintf(`{"data":"Link updated successfully to: %s","message":null,"success":true}`, "json")},
		{"", "https://stackoverflow.com/questions/33659298/unit-testing-http-json-response-in-golang", `{"data":null,"message":"ShortURL must be provided","success":false}`},
		{"asd", "", `{"data":null,"message":"LongURL must be provided","success":false}`},
	}

	for _, data := range testCase {
		m := new(mock.MockShortenerService)

		m.On("EditLink").Return(nil)

		router := httprouter.New()

		s := NewShortenerController(m)

		router.POST("/link/edit", s.EditLink())

		params := models.Link{
			LongUrl:  data.longUrl,
			ShortUrl: data.shortUrl,
		}

		j, err := json.Marshal(params)
		assert.Nil(t, err)

		req, _ := http.NewRequest("POST", "/link/edit", bytes.NewReader(j))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		// m.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, data.expected, rr.Body.String())
	}
}

func TestControllerDeleteLink(t *testing.T) {
	testCase := []struct {
		shortUrl string
		expected string
	}{
		{"", `{"data":null,"message":"ShortURL must be provided","success":false}`},
		{"asd", `{"data":"Link have been successfully deleted","message":null,"success":true}`},
	}

	for _, data := range testCase {

		m := new(mock.MockShortenerService)

		m.On("DeleteLink").Return(nil)

		router := httprouter.New()

		s := NewShortenerController(m)

		router.POST("/link/delete", s.DeleteLink())

		params := models.DeleteLink{
			ShortUrl: data.shortUrl,
		}

		j, err := json.Marshal(params)
		assert.Nil(t, err)

		req, _ := http.NewRequest("POST", "/link/delete", bytes.NewReader(j))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		// m.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, data.expected, rr.Body.String())
	}

}
