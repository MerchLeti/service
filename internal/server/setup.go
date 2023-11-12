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

	categoriesRepo := repository.Categories(db)
	itemsRepo := repository.Items(db)
	categories := endpoints.Categories(categoriesRepo)
	typesRepo := repository.Types(db)
	imagesRepo := repository.Images(db)
	properties := repository.Properties(db)
	itemsService := service.Items(itemsRepo, categoriesRepo, typesRepo, imagesRepo, properties)
	items := endpoints.Items(itemsService)

	api := r.PathPrefix("/api/").Subrouter()
	api.Handle("/categories", wrap(categories.GetAll)).Methods(http.MethodGet)
	api.Handle("/categories/{category}", wrap(idkeys.Category.UseIn(items.GetByCategory))).Methods(http.MethodGet)
	api.Handle("/items/{item}", wrap(idkeys.Item.UseIn(items.GetByID))).Methods(http.MethodGet)
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
