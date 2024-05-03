package mocks

import (
	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/x/auth/codec"
	"github.com/noble-assets/authority/x/authority/types"
)

var _ types.AccountKeeper = AccountKeeper{}

type AccountKeeper struct{}

func (AccountKeeper) AddressCodec() address.Codec {
	return codec.NewBech32Codec("noble")
}
