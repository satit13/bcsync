package app

import (
	"encoding/json"
	"errors"
	"github.com/acoshift/hrpc"
	//"github.com/mrtomyum/paybox_cloud/auth"
	"net/http"
	"fmt"
)

var (
	errMethodNotAllowed = errors.New("mobile: method not allowed")
	errForbidden = errors.New("mobile: forbidden")
)

type errorResponse struct {
	Error string `json:"error"`
}

// MakeHandler creates new mobile handler
func MakeHandler(s Service) http.Handler {
	m := hrpc.New(hrpc.Config{
		Validate:        true,
		RequestDecoder:  requestDecoder,
		ResponseEncoder: responseEncoder,
		ErrorEncoder:    errorEncoder,
	})

	mux := http.NewServeMux()
	mux.Handle("/create", m.Handler(makeAddInvoice(s)))
	return mustLogin()(mux)
}

func mustLogin() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//clientID := auth.GetClientID(r.Context())
			//if clientID < 0 {
			//	errorEncoder(w, r, errForbidden)
			//	return
			//}
			h.ServeHTTP(w, r)
		})
	}
}

func jsonDecoder(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func jsonEncoder(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func requestDecoder(r *http.Request, v interface{}) error {
	fmt.Println("r.Method = ", r.Method)
	if (r.Method == http.MethodPost || r.Method == http.MethodGet) {
		return errMethodNotAllowed
	}
	// TODO: choose decoder from request's content type
	// right now we have only json decoder
	return jsonDecoder(r, v)
}

func responseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) {
	// TODO: choose encoder from request's accept
	// right now we have only json encoder
	jsonEncoder(w, http.StatusOK, v)
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	// TODO: choose encoder from request's accept
	encoder := jsonEncoder

	//var status = http.StatusInternalServerError
	var status = http.StatusOK
	var response = errorResponse{err.Error()}

	switch err {
	case errMethodNotAllowed:
		status = http.StatusOK
	case errForbidden:
		status = http.StatusOK
	}

	encoder(w, status, &response)
}
