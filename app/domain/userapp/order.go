package userapp

import (
	"github.com/natnael-alemayehu/backend/business/domain/userbus"
)

var orderByFields = map[string]string{
	"user_id": userbus.OrderByID,
	"name":    userbus.OrderByName,
	"email":   userbus.OrderByEmail,
	"roles":   userbus.OrderByRoles,
	"enabled": userbus.OrderByEnabled,
}
