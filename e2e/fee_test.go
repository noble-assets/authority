// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package e2e

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"

	"github.com/stretchr/testify/require"
)

var COLLECTOR = "noble17xpfvakm2amg962yls6f84z3kell8c5lc6kgnn"

// TestBeginBlocker tests the module's begin blocker logic.
func TestBeginBlocker(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, _, _, _ := Suite(t, &wrapper, false)
	validator := wrapper.chain.Validators[0]

	oldBalance, err := wrapper.chain.BankQueryAllBalances(ctx, wrapper.owner.FormattedAddress())
	require.NoError(t, err)

	err = validator.BankSend(ctx, wrapper.owner.KeyName(), ibc.WalletAmount{
		Address: wrapper.pendingOwner.FormattedAddress(),
		Denom:   "uusdc",
		Amount:  math.NewInt(1_000_000),
	})
	require.NoError(t, err)

	balance, err := wrapper.chain.BankQueryAllBalances(ctx, COLLECTOR)
	require.NoError(t, err)
	require.True(t, balance.IsZero())

	newBalance, err := wrapper.chain.BankQueryAllBalances(ctx, wrapper.owner.FormattedAddress())
	require.NoError(t, err)
	require.Equal(t, oldBalance.Sub(sdk.NewInt64Coin("uusdc", 1_000_000)), newBalance)
}
