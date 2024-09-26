// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

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
