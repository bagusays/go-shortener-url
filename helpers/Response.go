package helpers

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, isSuccess bool, data interface{}, message interface{}) {
	payload := map[string]interface{}{
		"success": isSuccess,
		"data":    data,
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}
