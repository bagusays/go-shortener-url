package shortener

import (
	"encoding/json"
	"net/http"
	"url-shortener/domain"
	"url-shortener/helpers"
	"url-shortener/models"

	"github.com/julienschmidt/httprouter"
)

type ShortenerController struct {
	ShortenerService domain.ShortenerServiceInterface
}

func NewShortenerController(service domain.ShortenerServiceInterface) ShortenerController {
	return ShortenerController{ShortenerService: service}
}

func (s *ShortenerController) FindByShortLink() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		shortLink := p.ByName("shortLink")
		products, err := s.ShortenerService.FindByShortLink(shortLink)

		if err != nil {
			helpers.JSONResponse(w, false, nil, err.Error())
			return
		}

		if products == nil {
			helpers.JSONResponse(w, false, products, "404 Not Found")
			return
		}

		http.Redirect(w, r, products.LongUrl, 302)
	}
}

func (s *ShortenerController) CreateLink() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var createUrl models.Link
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&createUrl)
		if err != nil {
			helpers.JSONResponse(w, false, nil, "Something wrong")
			return
		}

		if createUrl.LongUrl == "" {
			helpers.JSONResponse(w, false, nil, "longUrl cannot be blank")
			return
		}

		if createUrl.ShortUrl == "" {
			helpers.JSONResponse(w, false, nil, "shortUrl cannot be blank")
			return
		}

		err = s.ShortenerService.CreateLink(createUrl)

		if err != nil {
			helpers.JSONResponse(w, false, nil, "Something wrong")
			return
		}

		helpers.JSONResponse(w, true, "http://localhost:8888/"+createUrl.ShortUrl, nil)
	}
}

func (s *ShortenerController) EditLink() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var link models.Link
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&link)
		if err != nil {
			helpers.JSONResponse(w, false, nil, "Something wrong")
			return
		}

		if link.LongUrl == "" {
			helpers.JSONResponse(w, false, nil, "longUrl cannot be blank")
			return
		}

		if link.ShortUrl == "" {
			helpers.JSONResponse(w, false, nil, "shortUrl cannot be blank")
			return
		}

		err = s.ShortenerService.EditLink(link)

		if err != nil {
			helpers.JSONResponse(w, false, nil, "Something wrong")
			return
		}

		helpers.JSONResponse(w, true, nil, "Link updated successfully to: "+link.ShortUrl)
	}
}

func (s *ShortenerController) DeleteLink() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var link models.DeleteLink
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&link)
		if err != nil {
			helpers.JSONResponse(w, false, nil, "Something wrong")
			return
		}
		if link.ShortUrl == "" {
			helpers.JSONResponse(w, false, nil, "shortUrl cannot be blank")
			return
		}

		err = s.ShortenerService.DeleteLink(link)

		if err != nil {
			helpers.JSONResponse(w, false, nil, "Something wrong")
			return
		}

		helpers.JSONResponse(w, true, nil, "Link updated successfully deleted")
	}
}
