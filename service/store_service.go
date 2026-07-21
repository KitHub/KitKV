package service

import (
	"context"
	"sync"

	"github.com/KitHub/kitkv/logic"
	"github.com/KitHub/protocols/kitkv"
)

var (
	storeServiceInstance *StoreService
	storeServiceOnce     sync.Once
)

type StoreService struct {
	kitkv.UnimplementedStoreServiceAPIServer
}

// Load implements [kitkv.StoreServiceAPIServer].
func (s *StoreService) Load(context.Context, *kitkv.LoadRequest) (*kitkv.LoadResponse, error) {
	panic("unimplemented")
}

// Store implements [kitkv.StoreServiceAPIServer].
func (s *StoreService) Store(context.Context, *kitkv.StoreRequest) (*kitkv.StoreResponse, error) {
	panic("unimplemented")
}

func NewStoreService(ctx context.Context, storeLogic *logic.StoreLogic) *StoreService {
	storeServiceOnce.Do(func() {
		storeServiceInstance = &StoreService{}
	})
	return storeServiceInstance
}
