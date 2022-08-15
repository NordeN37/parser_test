package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// RestResponse a common response for different routes if something went wrong
type RestResponse struct {
	Message string // Response to display to the client
	Details string // Answer details
}

const (
	MsgRequestError       = "ошибка обработки запроса"
	MsgJsonMarshalError   = "ошибка преобразования данных в json"
	MsgResponseWriteError = "не удалось создать ответ клиенту"
)

type HandlerFunc = func(w http.ResponseWriter, req *http.Request) (interface{}, error)

func wrapJSONHandler(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := handler(w, r)
		if err != nil {
			writeErrorResponseWeb(w, r, err)
			return
		}
		resBytes, err := json.Marshal(res)
		if err != nil {
			log.Printf("[ERROR] : %s , %s", MsgJsonMarshalError, err.Error())
			return
		}
		w.Header().Add("Content-Type", "application/json")
		func() {
			if _, err := w.Write(resBytes); err != nil {
				log.Printf("[ERROR] : %s , %s", MsgResponseWriteError, err.Error())
			}
		}()
	}
}

func writeErrorResponseWeb(w http.ResponseWriter, req *http.Request, err error) {
	log.Printf("[ERROR] : %s , %s", MsgRequestError, err.Error())

	restResponse := RestResponse{
		Message: MsgRequestError,
		Details: err.Error(),
	}
	var statusCode int
	statusCode, restResponse.Message = StatusCodeAndErrorMessage(err)

	resBytes, _ := json.Marshal(restResponse)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(resBytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func StatusCodeAndErrorMessage(err error) (int, string) {
	return http.StatusInternalServerError, errors.New(fmt.Sprintf("%s : %s", "Ошибка обработки запроса", err.Error())).Error()
}
