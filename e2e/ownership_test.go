// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package e2e

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestOwnershipTransfer tests a happy path two-step ownership transfer. This is
// to ensure that transactions are correctly registered, and that the CLI is
// generated correctly. Full test coverage is done via unit testing.
func TestOwnershipTransfer(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, _, _ := Suite(t, &wrapper, false)
	validator := wrapper.chain.Validators[0]

	EnsureOwner(t, wrapper, ctx, wrapper.owner.FormattedAddress())
	EnsurePendingOwner(t, wrapper, ctx, "")

	_, err := validator.ExecTx(
		ctx, wrapper.owner.KeyName(),
		"authority", "transfer-ownership", wrapper.pendingOwner.FormattedAddress(),
	)
	require.NoError(t, err)

	EnsureOwner(t, wrapper, ctx, wrapper.owner.FormattedAddress())
	EnsurePendingOwner(t, wrapper, ctx, wrapper.pendingOwner.FormattedAddress())

	_, err = validator.ExecTx(
		ctx, wrapper.pendingOwner.KeyName(),
		"authority", "accept-ownership",
	)
	require.NoError(t, err)

	EnsureOwner(t, wrapper, ctx, wrapper.pendingOwner.FormattedAddress())
	EnsurePendingOwner(t, wrapper, ctx, "")
}
