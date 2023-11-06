package endpoints

import (
	"net/http"

	"github.com/MerchLeti/service/internal/server/request"
)

type CategoriesEndpointsGroup struct {
	repo categoriesRepo
}

func Categories(repo categoriesRepo) *CategoriesEndpointsGroup {
	return &CategoriesEndpointsGroup{repo: repo}
}

func (c *CategoriesEndpointsGroup) GetAll(req *request.Request) {
	all, err := c.repo.GetAll(req.Context())
	if err != nil {
		req.WriteError(err)
		return
	}
	req.WriteObject(http.StatusOK, all)
}
