package endpoints

import (
	"net/http"

	"github.com/MerchLeti/service/internal/server/idkeys"
	"github.com/MerchLeti/service/internal/server/request"
)

const defaultCount = 10

type ItemsEndpointsGroup struct {
	repo itemsService
}

func Items(repo itemsService) *ItemsEndpointsGroup {
	return &ItemsEndpointsGroup{repo: repo}
}

func (c *ItemsEndpointsGroup) GetByCategory(req *request.Request) {
	page, err := req.GetQueryValueInt(false, "page", 1, 1)
	if err != nil {
		req.WriteErrorWithStatus(http.StatusBadRequest, err)
		return
	}
	count, err := req.GetQueryValueInt(false, "count", defaultCount, 1)
	if err != nil {
		req.WriteErrorWithStatus(http.StatusBadRequest, err)
		return
	}
	categoryID := idkeys.Category.Get(req.Context())
	items, err := c.repo.GetFromCategory(req.Context(), categoryID, page, count)
	if err != nil {
		req.WriteError(err)
		return
	}
	req.WriteObject(http.StatusOK, items)
}

func (c *ItemsEndpointsGroup) GetByID(req *request.Request) {
	itemID := idkeys.Item.Get(req.Context())
	item, err := c.repo.GetItem(req.Context(), itemID)
	if err != nil {
		req.WriteError(err)
		return
	}
	req.WriteObject(http.StatusOK, item)
}
