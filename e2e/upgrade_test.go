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
