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

	"github.com/stretchr/testify/require"
)

// TestOwnershipTransfer tests a happy path two-step ownership transfer. This is
// to ensure that transactions are correctly registered, and that the CLI is
// generated correctly. Full test coverage is done via unit testing.
func TestOwnershipTransfer(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, _, _, _ := Suite(t, &wrapper, false)
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
