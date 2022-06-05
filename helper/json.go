package helper

import (
	"encoding/json"
	"iteung-api/exception"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	exception.PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, result interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(result)
	exception.PanicIfError(err)
}
