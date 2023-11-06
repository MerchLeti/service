package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MerchLeti/service/internal/server/auth"
	"github.com/MerchLeti/service/internal/server/entities"
)

type method interface {
	Generate(data map[string]interface{}) (string, error)
	Validate(tokenString string) (map[string]interface{}, error)
}

var ErrAuthMethodUnavailable = errors.New("authentication method unavailable")

var authMethod method

func SetAuthMethod(newAuthMethod method) {
	authMethod = newAuthMethod
}

func OptionalAuth(endpoint http.Handler) http.Handler {
	return authHandler(
		endpoint, func(writer http.ResponseWriter, request *http.Request, err error) {
			endpoint.ServeHTTP(writer, request)
		},
	)
}

func Auth(endpoint http.Handler) http.Handler {
	return authHandler(
		endpoint, func(writer http.ResponseWriter, request *http.Request, err error) {
			if errors.Is(err, ErrAuthMethodUnavailable) {
				writeError(http.StatusInternalServerError, writer, err)
			} else {
				writeError(http.StatusUnauthorized, writer, err)
			}
		},
	)
}

func authHandler(endpoint http.Handler, onError func(http.ResponseWriter, *http.Request, error)) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			if authMethod == nil {
				onError(writer, request, ErrAuthMethodUnavailable)
				return
			}
			tokenString := request.Header["Token"]
			if tokenString == nil {
				onError(writer, request, auth.ErrUnauthorized)
				return
			}
			data, err := authMethod.Validate(tokenString[0])
			if err != nil {
				onError(writer, request, err)
				return
			}
			for key, value := range data {
				request = request.WithContext(context.WithValue(request.Context(), key, value))
			}
			endpoint.ServeHTTP(writer, request)
		},
	)
}

func writeError(status int, writer http.ResponseWriter, err error) {
	wrapped := entities.Result{
		Error: &entities.Error{
			Code:    status,
			Message: err.Error(),
		},
	}
	b, err := json.Marshal(wrapped)
	if err != nil {
		b = []byte(fmt.Sprintf("couldn't marshal response: %v", err))
		status = http.StatusInternalServerError
	}
	writer.WriteHeader(status)
	writer.Write(b)
}
