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

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/authority/utils"
	"github.com/noble-assets/authority/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestBeginBlock(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances: make(map[string]sdk.Coins),
	}
	keeper, ctx := mocks.AuthorityKeeperWithBank(bank)

	// ACT: Attempt to run begin blocker with empty fee collector.
	err := keeper.BeginBlock(ctx)
	// ASSERT: The action should've succeeded due to empty account.
	require.NoError(t, err)

	// ARRANGE: Give the fee collector some balance.
	bank.Balances["noble17xpfvakm2amg962yls6f84z3kell8c5lc6kgnn"] = sdk.NewCoins(
		sdk.NewInt64Coin("uusdc", 20_000),
	)

	// ACT: Attempt to run begin blocker with no owner set.
	err = keeper.BeginBlock(ctx)
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "failed to get owner from state")

	// ARRANGE: Set an invalid owner in state.
	require.NoError(t, keeper.Owner.Set(ctx, "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"))

	// ACT: Attempt to run begin blocker with invalid owner set.
	err = keeper.BeginBlock(ctx)
	// ASSERT: The action should've failed due to invalid owner set.
	require.ErrorContains(t, err, "failed to decode owner address")

	// ARRANGE: Generate an owner account and set in state.
	owner := utils.TestAccount()
	require.NoError(t, keeper.Owner.Set(ctx, owner.Address))

	// ACT: Attempt to run begin blocker.
	err = keeper.BeginBlock(ctx)
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
}
