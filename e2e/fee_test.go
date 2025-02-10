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
