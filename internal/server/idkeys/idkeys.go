package idkeys

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MerchLeti/service/internal/server/request"
	"github.com/gorilla/mux"
)

type Key string

var (
	Category Key = "category"
	Item     Key = "item"
)

func (c Key) Error() string {
	return fmt.Sprintf("%s id is absent or invalid", string(c))
}

func (c Key) Store(ctx context.Context, value int64) context.Context {
	return context.WithValue(ctx, c, value)
}

func (c Key) Get(ctx context.Context) int64 {
	value, _ := ctx.Value(c).(int64)
	// если преобразовать не получилось, value будет содержать 0, а паники не будет
	return value
}

func (c Key) UseIn(next func(req *request.Request)) func(req *request.Request) {
	return func(req *request.Request) {
		valueRaw, ok := mux.Vars(req.Request)[string(c)]
		if !ok {
			req.WriteErrorWithStatus(http.StatusBadRequest, c)
			return
		}
		value, err := strconv.ParseInt(valueRaw, 10, 64)
		if err != nil {
			req.WriteErrorWithStatus(http.StatusBadRequest, c)
			return
		}
		req.Request = req.Request.WithContext(c.Store(req.Context(), value))
		next(req)
	}
}
