package mocks

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/codec"
	"github.com/noble-assets/authority/x/authority/types"
)

var cdc = codec.NewBech32Codec("noble")

var _ types.BankKeeper = BankKeeper{}

type BankKeeper struct {
	Balances map[string]sdk.Coins
}

func (k BankKeeper) GetAllBalances(_ context.Context, bz sdk.AccAddress) sdk.Coins {
	address, _ := cdc.BytesToString(bz)
	return k.Balances[address]
}

func (k BankKeeper) SendCoins(_ context.Context, _, _ sdk.AccAddress, _ sdk.Coins) error {
	return nil
}
