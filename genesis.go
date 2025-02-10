// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, NASD Inc. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
