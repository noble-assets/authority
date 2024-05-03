package keeper_test

import (
	"testing"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/noble-assets/authority/utils"
	"github.com/noble-assets/authority/utils/mocks"
	"github.com/noble-assets/authority/x/authority/keeper"
	"github.com/noble-assets/authority/x/authority/types"
	"github.com/stretchr/testify/require"
)

const MODULE = "noble13am065qmk680w86wya4u9refhnssqwcvgs0sfk"

func TestExecute(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper(t)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Set an authority in state.
	authority := utils.TestAccount()
	require.NoError(t, k.Authority.Set(ctx, authority.Bytes))

	// ACT: Attempt to execute with invalid signer address.
	_, err := server.Execute(ctx, &types.MsgExecute{
		Signer: "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn",
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "failed to decode signer address")

	// ACT: Attempt to execute with invalid signer.
	_, err = server.Execute(ctx, &types.MsgExecute{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidAuthority.Error())

	// ACT: Attempt to execute with an invalid message signer address.
	msg := types.NewMsgExecute(authority.Address, []sdk.Msg{
		&banktypes.MsgUpdateParams{
			Authority: "",
			Params:    banktypes.Params{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to invalid message signer address.
	require.ErrorContains(t, err, "unable to extract signers")

	// ACT: Attempt to execute with an invalid message signer.
	msg = types.NewMsgExecute(authority.Address, []sdk.Msg{
		&banktypes.MsgUpdateParams{
			Authority: "noble10d07y265gmmuvt4z0w9aw880jnsr700jjpxdwa",
			Params:    banktypes.Params{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to invalid message signer.
	require.ErrorContains(t, err, types.ErrInvalidAuthority.Error())

	// ACT: Attempt to execute an invalid message.
	msg = types.NewMsgExecute(authority.Address, []sdk.Msg{
		&banktypes.MsgUpdateParams{
			Authority: MODULE,
			Params:    banktypes.Params{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to invalid message.
	require.ErrorContains(t, err, types.ErrInvalidMessage.Error())

	// ACT: Attempt to execute a valid message that fails.
	msg = types.NewMsgExecute(authority.Address, []sdk.Msg{
		&upgradetypes.MsgCancelUpgrade{
			Authority: MODULE,
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've failed due to failing message.
	require.ErrorContains(t, err, "failed to execute message")

	// ACT: Attempt to execute.
	msg = types.NewMsgExecute(authority.Address, []sdk.Msg{
		&upgradetypes.MsgSoftwareUpgrade{
			Authority: MODULE,
			Plan:      upgradetypes.Plan{},
		},
	})
	_, err = server.Execute(ctx, msg)
	// ASSERT: The action should've succeeded.
	require.Nil(t, err)
}

func TestUpdateAuthority(t *testing.T) {
	k, ctx := mocks.AuthorityKeeper(t)
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to update authority with no state.
	_, err := server.UpdateAuthority(ctx, &types.MsgUpdateAuthority{})
	// ASSERT: The action should've failed.
	require.ErrorContains(t, err, "unable to retrieve authority from state")

	// ARRANGE: Set an authority in state.
	authority := utils.TestAccount()
	require.NoError(t, k.Authority.Set(ctx, authority.Bytes))

	// ACT: Attempt to update authority with invalid signer address.
	_, err = server.UpdateAuthority(ctx, &types.MsgUpdateAuthority{
		Signer: "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn",
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "failed to decode signer address")

	// ACT: Attempt to update authority with invalid signer.
	_, err = server.UpdateAuthority(ctx, &types.MsgUpdateAuthority{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidAuthority.Error())

	// ACT: Attempt to update authority with invalid new authority address.
	_, err = server.UpdateAuthority(ctx, &types.MsgUpdateAuthority{
		Signer:       authority.Address,
		NewAuthority: "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn",
	})
	// ASSERT: The action should've failed due to invalid new authority address.
	require.ErrorContains(t, err, "failed to decode new authority address")

	// ACT: Attempt to update authority with same authority.
	_, err = server.UpdateAuthority(ctx, &types.MsgUpdateAuthority{
		Signer:       authority.Address,
		NewAuthority: authority.Address,
	})
	// ASSERT: The action should've failed due to same authority.
	require.ErrorContains(t, err, types.ErrSameAuthority.Error())

	// ARRANGE: Generate a new authority.
	newAuthority := utils.TestAccount()

	// ACT: Attempt to update authority.
	_, err = server.UpdateAuthority(ctx, &types.MsgUpdateAuthority{
		Signer:       authority.Address,
		NewAuthority: newAuthority.Address,
	})
	// ASSERT: The action should've succeeded, and updated authority.
	require.NoError(t, err)
	res, err := k.Authority.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, newAuthority.Bytes, res)
}
