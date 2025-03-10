// Package crud binds the crud domain set of routes into the specified app.
package crud

import (
	"time"

	"github.com/natnael-alemayehu/backend/app/domain/checkapp"
	"github.com/natnael-alemayehu/backend/app/domain/homeapp"
	"github.com/natnael-alemayehu/backend/app/domain/productapp"
	"github.com/natnael-alemayehu/backend/app/domain/tranapp"
	"github.com/natnael-alemayehu/backend/app/domain/userapp"
	"github.com/natnael-alemayehu/backend/app/sdk/mux"
	"github.com/natnael-alemayehu/backend/business/domain/homebus"
	"github.com/natnael-alemayehu/backend/business/domain/homebus/stores/homedb"
	"github.com/natnael-alemayehu/backend/business/domain/productbus"
	"github.com/natnael-alemayehu/backend/business/domain/productbus/stores/productdb"
	"github.com/natnael-alemayehu/backend/business/domain/userbus"
	"github.com/natnael-alemayehu/backend/business/domain/userbus/stores/usercache"
	"github.com/natnael-alemayehu/backend/business/domain/userbus/stores/userdb"
	"github.com/natnael-alemayehu/backend/business/sdk/delegate"
	"github.com/natnael-alemayehu/backend/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	delegate := delegate.New(cfg.Log)
	userBus := userbus.NewBusiness(cfg.Log, delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB), time.Minute))
	productBus := productbus.NewBusiness(cfg.Log, userBus, delegate, productdb.NewStore(cfg.Log, cfg.DB))
	homeBus := homebus.NewBusiness(cfg.Log, userBus, delegate, homedb.NewStore(cfg.Log, cfg.DB))

	checkapp.Routes(app, checkapp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	homeapp.Routes(app, homeapp.Config{
		HomeBus:    homeBus,
		AuthClient: cfg.AuthClient,
	})

	productapp.Routes(app, productapp.Config{
		ProductBus: productBus,
		AuthClient: cfg.AuthClient,
	})

	tranapp.Routes(app, tranapp.Config{
		UserBus:    userBus,
		ProductBus: productBus,
		Log:        cfg.Log,
		AuthClient: cfg.AuthClient,
		DB:         cfg.DB,
	})

	userapp.Routes(app, userapp.Config{
		UserBus:    userBus,
		AuthClient: cfg.AuthClient,
	})
}
