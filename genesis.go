// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package authority

import (
	"context"

	"github.com/noble-assets/authority/keeper"
	"github.com/noble-assets/authority/types"
)

func InitGenesis(ctx context.Context, k *keeper.Keeper, accountKeeper types.AccountKeeper, genesis types.GenesisState) {
	_, err := accountKeeper.AddressCodec().StringToBytes(genesis.Owner)
	if err != nil {
		panic("failed to decode owner address")
	}
	err = k.Owner.Set(ctx, genesis.Owner)
	if err != nil {
		panic("failed to set owner in state")
	}

	if genesis.PendingOwner != "" {
		_, err := accountKeeper.AddressCodec().StringToBytes(genesis.PendingOwner)
		if err != nil {
			panic("failed to decode pending owner address")
		}
		err = k.PendingOwner.Set(ctx, genesis.PendingOwner)
		if err != nil {
			panic("failed to set pending owner in state")
		}
	}
}

func ExportGenesis(ctx context.Context, k *keeper.Keeper) *types.GenesisState {
	owner, _ := k.Owner.Get(ctx)
	pendingOwner, _ := k.PendingOwner.Get(ctx)

	return &types.GenesisState{
		Owner:        owner,
		PendingOwner: pendingOwner,
	}
}
