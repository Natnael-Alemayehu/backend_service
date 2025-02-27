package dbtest

import (
	"time"

	"github.com/natnael-alemayehu/backend/business/domain/homebus"
	"github.com/natnael-alemayehu/backend/business/domain/homebus/stores/homedb"
	"github.com/natnael-alemayehu/backend/business/domain/productbus"
	"github.com/natnael-alemayehu/backend/business/domain/productbus/stores/productdb"
	"github.com/natnael-alemayehu/backend/business/domain/userbus"
	"github.com/natnael-alemayehu/backend/business/domain/userbus/stores/usercache"
	"github.com/natnael-alemayehu/backend/business/domain/userbus/stores/userdb"
	"github.com/natnael-alemayehu/backend/business/domain/vproductbus"
	"github.com/natnael-alemayehu/backend/business/domain/vproductbus/stores/vproductdb"
	"github.com/natnael-alemayehu/backend/business/sdk/delegate"
	"github.com/natnael-alemayehu/backend/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	Delegate *delegate.Delegate
	Home     *homebus.Business
	Product  *productbus.Business
	User     *userbus.Business
	VProduct *vproductbus.Business
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	delegate := delegate.New(log)
	userBus := userbus.NewBusiness(log, delegate, usercache.NewStore(log, userdb.NewStore(log, db), time.Hour))
	productBus := productbus.NewBusiness(log, userBus, delegate, productdb.NewStore(log, db))
	homeBus := homebus.NewBusiness(log, userBus, delegate, homedb.NewStore(log, db))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(log, db))

	return BusDomain{
		Delegate: delegate,
		Home:     homeBus,
		Product:  productBus,
		User:     userBus,
		VProduct: vproductBus,
	}
}
