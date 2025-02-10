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
	_ "embed"
	"testing"
	"time"

	"cosmossdk.io/math"

	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
)

// TestClientSubstitution tests the module's ability to substitute IBC clients.
func TestClientSubstitution(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, logger, reporter, rly := Suite(t, &wrapper, true)
	validator := wrapper.chain.Validators[0]

	nobleChainID, gaiaChainID := wrapper.chain.Config().ChainID, wrapper.gaia.Config().ChainID
	pathName := "transfer"

	err := rly.GeneratePath(ctx, reporter, nobleChainID, gaiaChainID, pathName)
	require.NoError(t, err)

	tp := 20 * time.Second
	err = rly.CreateClient(ctx, reporter, nobleChainID, gaiaChainID, pathName, ibc.CreateClientOptions{
		TrustingPeriod: tp.String(),
	})
	require.NoError(t, err)

	nobleClients, err := rly.GetClients(ctx, reporter, nobleChainID)
	require.NoError(t, err)
	require.Len(t, nobleClients, 2) // ignore 09-localhost client

	nobleClientToExpire := nobleClients[0]

	err = rly.CreateClient(ctx, reporter, gaiaChainID, nobleChainID, pathName, ibc.CreateClientOptions{})
	require.NoError(t, err)
	require.NoError(t, testutil.WaitForBlocks(ctx, 1, wrapper.chain, wrapper.gaia))

	err = rly.CreateConnections(ctx, reporter, pathName)
	require.NoError(t, err)
	require.NoError(t, testutil.WaitForBlocks(ctx, 1, wrapper.chain, wrapper.gaia))

	err = rly.CreateChannel(ctx, reporter, pathName, ibc.DefaultChannelOpts())
	require.NoError(t, err)

	timer := time.NewTimer(tp + 2*time.Second)

	users := interchaintest.GetAndFundTestUsers(t, ctx, "user", math.NewInt(5_000_000), wrapper.chain, wrapper.gaia)

	logger.Info("waiting for client to expire...")
	<-timer.C

	_, err = validator.SendIBCTransfer(ctx, "channel-0", users[0].KeyName(), ibc.WalletAmount{
		Address: users[1].FormattedAddress(),
		Denom:   "uusdc",
		Amount:  math.NewInt(1_000_000),
	}, ibc.TransferOptions{})
	require.ErrorContains(t, err, "client state is not active")

	err = rly.CreateClient(ctx, reporter, nobleChainID, gaiaChainID, pathName, ibc.CreateClientOptions{Override: true})
	require.NoError(t, err)

	nobleClients, err = rly.GetClients(ctx, reporter, nobleChainID)
	require.NoError(t, err)
	require.Len(t, nobleClients, 3) // ignore 09-localhost client

	newNobleClient := nobleClients[1]

	notAuthorized := interchaintest.GetAndFundTestUsers(t, ctx, "wallet", math.NewInt(100000), wrapper.chain)[0]

	cmd := []string{"authority", "recover-client", nobleClientToExpire.ClientID, newNobleClient.ClientID}

	// broadcast from un-authorized account
	_, err = validator.ExecTx(
		ctx,
		notAuthorized.KeyName(),
		cmd...,
	)
	require.ErrorContains(t, err, "signer is not authority")

	// broadcast from authorized authority account
	_, err = validator.ExecTx(
		ctx,
		wrapper.owner.KeyName(),
		cmd...,
	)
	require.NoError(t, err)

	_, err = validator.SendIBCTransfer(ctx, "channel-0", users[0].KeyName(), ibc.WalletAmount{
		Address: users[1].FormattedAddress(),
		Denom:   "uusdc",
		Amount:  math.NewInt(1_000_000),
	}, ibc.TransferOptions{})
	require.NoError(t, err)
}
