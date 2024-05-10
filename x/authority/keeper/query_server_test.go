package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/authority/utils"
	"github.com/noble-assets/authority/utils/mocks"
	"github.com/noble-assets/authority/x/authority/keeper"
	"github.com/noble-assets/authority/x/authority/types"
	"github.com/stretchr/testify/require"
)

func TestAddressQuery(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper(t)
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query address with invalid request.
	_, err := server.Address(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query address with no state.
	_, err = server.Address(ctx, &types.QueryAddress{})
	// ASSERT: The query should've failed.
	require.ErrorContains(t, err, "unable to retrieve authority from state")

	// ARRANGE: Set an authority in state.
	authority := utils.TestAccount()
	require.NoError(t, k.Authority.Set(ctx, authority.Bytes))

	// ACT: Attempt to query address.
	res, err := server.Address(ctx, &types.QueryAddress{})
	// ASSERT: The query should've succeeded, and returned address.
	require.NoError(t, err)
	require.Equal(t, authority.Address, res.Address)

	// ARRANGE: Set a pending authority in state.
	pendingAuthority := utils.TestAccount()
	require.NoError(t, k.PendingAuthority.Set(ctx, pendingAuthority.Bytes))

	// ACT: Attempt to query address.
	res, err = server.Address(ctx, &types.QueryAddress{})
	// ASSERT: The query should've succeeded, and returned address and pending address.
	require.NoError(t, err)
	require.Equal(t, authority.Address, res.Address)
	require.Equal(t, pendingAuthority.Address, res.PendingAddress)
}
