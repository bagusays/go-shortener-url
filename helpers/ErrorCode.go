package helpers

import "fmt"

type ErrorStruct struct {
	Code        string
	Description string
}

func ErrorList(err error) error {
	errorList := []ErrorStruct{
		{Code: "SHORT_URL_IS_TAKEN", Description: "ShortUrl is taken for another url. Please select another shortUrl"},
		{Code: "NOT_FOUND", Description: "404 Not Found"},
		{Code: "MISSING_SHORTURL", Description: "ShortURL must be provided"},
		{Code: "MISSING_LONGURL", Description: "LongURL must be provided"},
		{Code: "NOT_FOUND", Description: "404 Not Found"},
	}

	if err == nil {
		return nil
	}

	for _, e := range errorList {
		if e.Code == err.Error() {
			return fmt.Errorf(e.Description)
		}
	}

	return fmt.Errorf("Something wrong")
}
