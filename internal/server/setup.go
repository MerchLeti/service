package server

import (
	"log"
	"net/http"
	"os"

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

	r.Handle("/api/categories", wrap(categories.GetAll)).Methods(http.MethodGet)
	r.Handle("/api/categories/{category}", wrap(idkeys.Category.UseIn(items.GetByCategory))).Methods(http.MethodGet)
	r.Handle("/api/items/{item}", wrap(idkeys.Item.UseIn(items.GetByID))).Methods(http.MethodGet)
	if frontendExists() {
		log.Println("Including frontend file server")
		r.PathPrefix("/").Handler(http.FileServer(http.Dir("/frontend/")))
	}
	return r
}

func frontendExists() bool {
	stat, err := os.Stat("/frontend")
	return err == nil && stat.IsDir()
}

func wrap(nxt func(req *request.Request)) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, r *http.Request) {
			nxt(request.New(writer, r))
		},
	)
}
