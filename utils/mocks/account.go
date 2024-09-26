// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package mocks

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/noble-assets/authority/types"
)

var _ types.AccountKeeper = AccountKeeper{}

type AccountKeeper struct{}

func (AccountKeeper) AddressCodec() address.Codec {
	return codec.NewBech32Codec("noble")
}

func (k AccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return authtypes.NewModuleAddress(moduleName)
}
