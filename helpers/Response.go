package helpers

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, isSuccess bool, data interface{}, message error) {
	var err interface{}

	if isSuccess == false {
		message = ErrorList(message)
		err = message.Error()
	} else {
		err = nil
	}

	payload := map[string]interface{}{
		"success": isSuccess,
		"data":    data,
		"message": err,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}
