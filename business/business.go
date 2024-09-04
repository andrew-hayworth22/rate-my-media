package business

import (
	"context"

	"github.com/andrew-hayworth22/rate-my-media/database"
)

type BusinessLayer interface {
	CreateUser(context.Context, BusStoreUserRequest) (BusUser, error)
}

type Business struct {
	db database.DatabaseLayer
}

func NewBusiness(db database.DatabaseLayer) Business {
	return Business{
		db: db,
	}
}
