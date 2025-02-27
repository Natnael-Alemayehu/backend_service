// Package vproductapp maintains the app layer api for the vproduct domain.
package vproductapp

import (
	"context"
	"net/http"

	"github.com/natnael-alemayehu/backend/app/sdk/errs"
	"github.com/natnael-alemayehu/backend/app/sdk/query"
	"github.com/natnael-alemayehu/backend/business/domain/vproductbus"
	"github.com/natnael-alemayehu/backend/business/sdk/order"
	"github.com/natnael-alemayehu/backend/business/sdk/page"
	"github.com/natnael-alemayehu/backend/foundation/web"
)

type app struct {
	vproductBus *vproductbus.Business
}

func newApp(vproductBus *vproductbus.Business) *app {
	return &app{
		vproductBus: vproductBus,
	}
}

func (a *app) query(ctx context.Context, r *http.Request) web.Encoder {
	qp := parseQueryParams(r)

	page, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return errs.NewFieldErrors("page", err)
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return err.(*errs.Error)
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, vproductbus.DefaultOrderBy)
	if err != nil {
		return errs.NewFieldErrors("order", err)
	}

	prds, err := a.vproductBus.Query(ctx, filter, orderBy, page)
	if err != nil {
		return errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.vproductBus.Count(ctx, filter)
	if err != nil {
		return errs.Newf(errs.Internal, "count: %s", err)
	}

	return query.NewResult(toAppProducts(prds), total, page)
}
