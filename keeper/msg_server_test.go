// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"testing"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/noble-assets/authority/keeper"
	"github.com/noble-assets/authority/types"
	"github.com/noble-assets/authority/utils"
	"github.com/noble-assets/authority/utils/mocks"
	"github.com/stretchr/testify/require"
)

const MODULE = "noble13am065qmk680w86wya4u9refhnssqwcvgs0sfk"

func TestExecute(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper()
	server := keeper.NewMsgServer(k)

	// ARRANGE: Set an owner in state.
	owner := utils.TestAccount()
	require.NoError(t, k.Owner.Set(ctx, owner.Address))

	// ACT: Attempt to execute with invalid signer.
	_, err := server.Execute(ctx, &types.MsgExecute{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidOwner.Error())

	// ACT: Attempt to execute with an invalid message signer address.
	msg := types.NewMsgExecute(owner.Address, []sdk.Msg{
		&banktypes.MsgUpdateParams{
			Authority: "",
			Params:    banktypes.Params{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to invalid message signer address.
	require.ErrorContains(t, err, "unable to extract signers")

	// ACT: Attempt to execute with an invalid message signer.
	msg = types.NewMsgExecute(owner.Address, []sdk.Msg{
		&banktypes.MsgUpdateParams{
			Authority: "noble10d07y265gmmuvt4z0w9aw880jnsr700jjpxdwa",
			Params:    banktypes.Params{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to invalid message signer.
	require.ErrorContains(t, err, types.ErrInvalidOwner.Error())

	// ACT: Attempt to execute an invalid message.
	msg = types.NewMsgExecute(owner.Address, []sdk.Msg{
		&banktypes.MsgUpdateParams{
			Authority: MODULE,
			Params:    banktypes.Params{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to invalid message.
	require.ErrorContains(t, err, types.ErrInvalidMessage.Error())

	// ACT: Attempt to execute a valid message that fails.
	msg = types.NewMsgExecute(owner.Address, []sdk.Msg{
		&upgradetypes.MsgCancelUpgrade{
			Authority: MODULE,
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to failing message.
	require.ErrorContains(t, err, "failed to execute message")

	// ACT: Attempt to execute.
	msg = types.NewMsgExecute(owner.Address, []sdk.Msg{
		&upgradetypes.MsgSoftwareUpgrade{
			Authority: MODULE,
			Plan:      upgradetypes.Plan{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've succeeded.
	require.Nil(t, err)
}

func TestTransferOwnership(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper()
	server := keeper.NewMsgServer(k)

	// ARRANGE: Set an owner in state.
	owner := utils.TestAccount()
	require.NoError(t, k.Owner.Set(ctx, owner.Address))

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err := server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership with invalid new owner address.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn",
	})
	// ASSERT: The action should've failed due to invalid new owner address.
	require.ErrorContains(t, err, "failed to decode new owner address")

	// ACT: Attempt to transfer ownership with same owner.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same owner.
	require.ErrorContains(t, err, types.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner.
	newOwner := utils.TestAccount()

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set a pending owner.
	require.NoError(t, err)
	res, err := k.Owner.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, owner.Address, res)
	res, err = k.PendingOwner.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, res)
}

func TestAcceptOwnership(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper()
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to accept ownership with no pending owner set.
	_, err := server.AcceptOwnership(ctx, &types.MsgAcceptOwnership{})
	// ASSERT: The action should've failed.
	require.ErrorIs(t, err, types.ErrNoPendingOwner)

	// ARRANGE: Set a pending owner in state.
	pendingOwner := utils.TestAccount()
	require.NoError(t, k.PendingOwner.Set(ctx, pendingOwner.Address))

	// ACT: Attempt to accept ownership with invalid signer.
	_, err = server.AcceptOwnership(ctx, &types.MsgAcceptOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidPendingOwner.Error())

	// ACT: Attempt to accept ownership.
	_, err = server.AcceptOwnership(ctx, &types.MsgAcceptOwnership{
		Signer: pendingOwner.Address,
	})
	// ASSERT: The action should've succeeded, and updated owner.
	require.NoError(t, err)
	has, err := k.PendingOwner.Has(ctx)
	require.NoError(t, err)
	require.False(t, has)
	res, err := k.Owner.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, pendingOwner.Address, res)
}
