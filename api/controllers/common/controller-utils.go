package utils

import (
	"encoding/json"
	"net/http"
)

func WriteOkResponse[T any](object *T, writer http.ResponseWriter) {
	response, _ := json.Marshal(object)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
