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

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/authority/keeper"
	"github.com/noble-assets/authority/types"
	"github.com/noble-assets/authority/utils"
	"github.com/noble-assets/authority/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestOwnerQuery(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper()
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query owner with invalid request.
	_, err := server.Owner(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query owner with no state.
	_, err = server.Owner(ctx, &types.QueryOwner{})
	// ASSERT: The query should've failed.
	require.ErrorIs(t, err, collections.ErrNotFound)

	// ARRANGE: Set an owner in state.
	owner := utils.TestAccount()
	require.NoError(t, k.Owner.Set(ctx, owner.Address))

	// ACT: Attempt to query owner.
	res, err := server.Owner(ctx, &types.QueryOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
}

func TestPendingOwnerQuery(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper()
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query pending owner with invalid request.
	_, err := server.PendingOwner(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query pending owner with no state.
	res, err := server.PendingOwner(ctx, &types.QueryPendingOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Empty(t, res.PendingOwner)

	// ARRANGE: Set a pending owner in state.
	pendingOwner := utils.TestAccount()
	require.NoError(t, k.PendingOwner.Set(ctx, pendingOwner.Address))

	// ACT: Attempt to query pending owner.
	res, err = server.PendingOwner(ctx, &types.QueryPendingOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, pendingOwner.Address, res.PendingOwner)
}
