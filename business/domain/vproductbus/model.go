package vproductbus

import (
	"time"

	"github.com/natnael-alemayehu/backend/business/types/money"
	"github.com/natnael-alemayehu/backend/business/types/name"
	"github.com/natnael-alemayehu/backend/business/types/quantity"
	"github.com/google/uuid"
)

// Product represents an individual product with extended information.
type Product struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        name.Name
	Cost        money.Money
	Quantity    quantity.Quantity
	DateCreated time.Time
	DateUpdated time.Time
	UserName    name.Name
}
