package server

import (
	"log"
	"net/http"

	"github.com/MerchLeti/service/internal/repository"
	"github.com/MerchLeti/service/internal/server/endpoints"
	"github.com/MerchLeti/service/internal/server/idkeys"
	"github.com/MerchLeti/service/internal/server/request"
	"github.com/MerchLeti/service/internal/service"
	"github.com/gorilla/mux"
)

func New(db repository.DataSource) *mux.Router {
	r := mux.NewRouter()

	categories := repository.Categories(db)
	items := repository.Items(db)
	types := repository.Types(db)
	images := repository.Images(db)
	properties := repository.Properties(db)
	descriptions := repository.Descriptions(db)
	itemsService := service.Items(items, categories, types, images, properties, descriptions)
	itemsEndpoint := endpoints.Items(itemsService)

	api := r.PathPrefix("/api/").Subrouter()
	api.Handle("/categories", wrap(endpoints.Categories(categories).GetAll)).Methods(http.MethodGet)
	api.Handle(
		"/categories/{category}",
		wrap(idkeys.Category.UseIn(itemsEndpoint.GetByCategory)),
	).Methods(http.MethodGet)
	api.Handle("/items/{item}", wrap(idkeys.Item.UseIn(itemsEndpoint.GetByID))).Methods(http.MethodGet)
	if frontend := endpoints.Frontend(); frontend != nil {
		log.Println("Including frontend file server")
		r.PathPrefix("/").Handler(frontend)
	}
	return r
}

func wrap(nxt func(req *request.Request)) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, r *http.Request) {
			nxt(request.New(writer, r))
		},
	)
}
