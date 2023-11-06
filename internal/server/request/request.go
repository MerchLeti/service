package request

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/MerchLeti/service/internal/repository"
	"github.com/MerchLeti/service/internal/server/entities"
)

type Request struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func New(writer http.ResponseWriter, request *http.Request) *Request {
	return &Request{
		Writer:  writer,
		Request: request,
	}
}

func (r *Request) Context() context.Context {
	return r.Request.Context()
}

func (r *Request) WriteObject(status int, obj interface{}) {
	r.write(status, &entities.Result{Result: obj})
}

func (r *Request) WriteErrorWithStatus(status int, err error) {
	r.write(status, &entities.Result{Error: &entities.Error{Code: status, Message: err.Error()}})
}

func (r *Request) WriteError(err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		r.WriteErrorWithStatus(http.StatusNotFound, err)
	default:
		r.WriteErrorWithStatus(http.StatusInternalServerError, err)
	}
}

func (r *Request) write(status int, wrapped *entities.Result) {
	b, err := json.Marshal(wrapped)
	useJson := true
	if err != nil {
		b = []byte(fmt.Sprintf("couldn't marshal response to json: %v", err))
		status = http.StatusInternalServerError
		useJson = false
	}
	r.Writer.WriteHeader(status)
	if useJson {
		r.Writer.Header().Add("Content-Type", "application/json")
	}
	if _, err := r.Writer.Write(b); err != nil {
		log.Printf("couldn't write response: %v\n", err)
	}
}
