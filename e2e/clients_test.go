// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package e2e

import (
	_ "embed"
	"path"
	"testing"

	"cosmossdk.io/math"

	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
)

//go:embed clients.json
var Clients []byte

// TestClientSubstitution tests the module's ability to substitute IBC clients.
func TestClientSubstitution(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, reporter, rly := Suite(t, &wrapper, true)
	validator := wrapper.chain.Validators[0]

	nobleChainID, gaiaChainID := wrapper.chain.Config().ChainID, wrapper.gaia.Config().ChainID
	pathName := "transfer"

	err := rly.GeneratePath(ctx, reporter, nobleChainID, gaiaChainID, pathName)
	require.NoError(t, err)

	err = rly.CreateClient(ctx, reporter, nobleChainID, gaiaChainID, pathName, ibc.CreateClientOptions{
		TrustingPeriod: "30s",
	})
	require.NoError(t, err)
	err = rly.CreateClient(ctx, reporter, gaiaChainID, nobleChainID, pathName, ibc.CreateClientOptions{})
	require.NoError(t, err)
	require.NoError(t, testutil.WaitForBlocks(ctx, 1, wrapper.chain, wrapper.gaia))

	err = rly.CreateConnections(ctx, reporter, pathName)
	require.NoError(t, err)
	require.NoError(t, testutil.WaitForBlocks(ctx, 1, wrapper.chain, wrapper.gaia))

	err = rly.CreateChannel(ctx, reporter, pathName, ibc.DefaultChannelOpts())
	require.NoError(t, err)

	users := interchaintest.GetAndFundTestUsers(t, ctx, "user", math.NewInt(5_000_000), wrapper.chain, wrapper.gaia)
	require.NoError(t, testutil.WaitForBlocks(ctx, 10, wrapper.chain, wrapper.gaia))

	_, err = validator.SendIBCTransfer(ctx, "channel-0", users[0].KeyName(), ibc.WalletAmount{
		Address: users[1].FormattedAddress(),
		Denom:   "uusdc",
		Amount:  math.NewInt(1_000_000),
	}, ibc.TransferOptions{})
	require.ErrorContains(t, err, "client state is not active")

	res := rly.Exec(ctx, reporter, []string{"rly", "tx", "client", nobleChainID, gaiaChainID, pathName, "--override", "--home", rly.HomeDir()}, nil)
	require.NoError(t, res.Err)

	require.NoError(t, validator.WriteFile(ctx, Clients, "clients.json"))
	_, err = validator.ExecTx(
		ctx, wrapper.owner.KeyName(),
		"authority", "execute", path.Join(validator.HomeDir(), "clients.json"),
	)
	require.NoError(t, err)

	_, err = validator.SendIBCTransfer(ctx, "channel-0", users[0].KeyName(), ibc.WalletAmount{
		Address: users[1].FormattedAddress(),
		Denom:   "uusdc",
		Amount:  math.NewInt(1_000_000),
	}, ibc.TransferOptions{})
	require.NoError(t, err)
}
