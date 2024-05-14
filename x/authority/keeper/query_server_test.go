package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/authority/utils"
	"github.com/noble-assets/authority/utils/mocks"
	"github.com/noble-assets/authority/x/authority/keeper"
	"github.com/noble-assets/authority/x/authority/types"
	"github.com/stretchr/testify/require"
)

func TestOwnerQuery(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper(t)
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
	k, ctx := mocks.AuthorityKeeper(t)
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
