package shortener

import (
	"fmt"
	"url-shortener/domain"
	"url-shortener/models"
)

type ShortenerService struct {
	ShortenerRepo domain.ShortenerRepoInterface
}

func NewShortenerService(repo domain.ShortenerRepoInterface) *ShortenerService {
	return &ShortenerService{ShortenerRepo: repo}
}

func (s *ShortenerService) FindByShortLink(shortLink string) (*models.Link, error) {
	products, err := s.ShortenerRepo.FindByShortLink(shortLink)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *ShortenerService) CreateLink(data models.Link) error {
	checkShortUrl, _ := s.FindByShortLink(data.ShortUrl)
	if checkShortUrl != nil {
		return fmt.Errorf("SHORT_URL_IS_TAKEN")
	}

	err := s.ShortenerRepo.CreateLink(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShortenerService) EditLink(data models.Link) error {
	err := s.ShortenerRepo.EditLink(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShortenerService) DeleteLink(data models.DeleteLink) error {
	err := s.ShortenerRepo.DeleteLink(data)
	if err != nil {
		return err
	}
	return nil
}
