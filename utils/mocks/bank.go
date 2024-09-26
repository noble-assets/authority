// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package mocks

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/codec"
	"github.com/noble-assets/authority/types"
)

var _ types.BankKeeper = BankKeeper{}

type BankKeeper struct {
	Balances map[string]sdk.Coins
}

func (k BankKeeper) GetAllBalances(_ context.Context, bz sdk.AccAddress) sdk.Coins {
	address, _ := codec.NewBech32Codec("noble").BytesToString(bz)
	return k.Balances[address]
}

func (k BankKeeper) SendCoins(_ context.Context, _, _ sdk.AccAddress, _ sdk.Coins) error {
	return nil
}
