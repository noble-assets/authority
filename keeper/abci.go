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

package keeper

import (
	"context"

	"cosmossdk.io/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// BeginBlock sends all fees collected in the previous block to the module's owner.
func (k *Keeper) BeginBlock(ctx context.Context) error {
	collector := k.accountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	balance := k.bankKeeper.GetAllBalances(ctx, collector)
	if balance.IsZero() {
		return nil
	}

	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get owner from state")
	}
	address, err := k.accountKeeper.AddressCodec().StringToBytes(owner)
	if err != nil {
		return errors.Wrap(err, "failed to decode owner address")
	}

	return k.bankKeeper.SendCoins(ctx, collector, address, balance)
}
