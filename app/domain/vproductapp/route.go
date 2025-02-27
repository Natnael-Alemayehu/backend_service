package vproductapp

import (
	"net/http"

	"github.com/natnael-alemayehu/backend/app/sdk/auth"
	"github.com/natnael-alemayehu/backend/app/sdk/authclient"
	"github.com/natnael-alemayehu/backend/app/sdk/mid"
	"github.com/natnael-alemayehu/backend/business/domain/userbus"
	"github.com/natnael-alemayehu/backend/business/domain/vproductbus"
	"github.com/natnael-alemayehu/backend/foundation/logger"
	"github.com/natnael-alemayehu/backend/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	UserBus     *userbus.Business
	VProductBus *vproductbus.Business
	AuthClient  *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newApp(cfg.VProductBus)

	app.HandlerFunc(http.MethodGet, version, "/vproducts", api.query, authen, ruleAdmin)
}
