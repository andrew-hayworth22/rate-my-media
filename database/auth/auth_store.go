package auth

import (
	"context"
)

type Store interface {
	StoreUser(ctx context.Context, req DbStoreUserRequest) (DbUser, error)
	GetUserByEmail(ctx context.Context, email string) (DbUser, error)
}
