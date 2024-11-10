// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package e2e

import (
	_ "embed"
	"testing"

	"cosmossdk.io/math"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/stretchr/testify/require"
)

// TestScheduleUpgrade tests the module's ability to schedule an upgrade on-chain.
func TestScheduleUpgrade(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, _, _, _ := Suite(t, &wrapper, false)
	validator := wrapper.chain.Validators[0]

	notAuthorized := interchaintest.GetAndFundTestUsers(t, ctx, "wallet", math.NewInt(100000), wrapper.chain)[0]

	EnsureUpgrade(t, wrapper, ctx, "", 0)

	cmd := []string{"authority", "software-upgrade", "v2", "--upgrade-height", "50",
		"--chain-id", wrapper.chain.Config().ChainID, "--no-validate"}

	// broadcast from un-authorized account
	_, err := validator.ExecTx(
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

	EnsureUpgrade(t, wrapper, ctx, "v2", 50)
}
