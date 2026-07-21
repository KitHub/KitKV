package logic

import (
	"context"
	"sync"
)

var storeLogic *StoreLogic
var onceStoreLogic sync.Once

type StoreLogic struct {
}

// Hello implements [DemoLogic].
func (d *StoreLogic) Hello(ctx context.Context) error {
	panic("unimplemented")
}

func NewStoreLogic(ctx context.Context) *StoreLogic {
	onceStoreLogic.Do(func() {
		storeLogic = &StoreLogic{}
	})
	return storeLogic
}
