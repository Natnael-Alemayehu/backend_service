package tranapp

import (
	"net/http"

	"github.com/natnael-alemayehu/backend/app/sdk/auth"
	"github.com/natnael-alemayehu/backend/app/sdk/authclient"
	"github.com/natnael-alemayehu/backend/app/sdk/mid"
	"github.com/natnael-alemayehu/backend/business/domain/productbus"
	"github.com/natnael-alemayehu/backend/business/domain/userbus"
	"github.com/natnael-alemayehu/backend/business/sdk/sqldb"
	"github.com/natnael-alemayehu/backend/foundation/logger"
	"github.com/natnael-alemayehu/backend/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	DB         *sqlx.DB
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	transaction := mid.BeginCommitRollback(cfg.Log, sqldb.NewBeginner(cfg.DB))
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newApp(cfg.UserBus, cfg.ProductBus)

	app.HandlerFunc(http.MethodPost, version, "/tranexample", api.create, authen, ruleAdmin, transaction)
}
