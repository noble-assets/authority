package keeper

import (
	"bytes"
	"context"

	"cosmossdk.io/errors"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/authority/x/authority/types"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Execute(ctx context.Context, msg *types.MsgExecute) (*types.MsgExecuteResponse, error) {
	_, err := k.EnsureAuthoritySigner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	msgs, err := msg.GetMessages(k.cdc)
	if err != nil {
		return nil, err
	}
	if err := validateMsgs(msgs); err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var events []abcitypes.Event
	var results [][]byte

	for i, msg := range msgs {
		// NOTE: Ensure that for each message, there is only one signer, and
		//  that that signer is the x/authority module.
		signers, _, err := k.cdc.GetMsgV1Signers(msg)
		if err != nil {
			return nil, errors.Wrapf(err, "message %d; unable to extract signers", i)
		}
		if len(signers) != 1 || !bytes.Equal(types.ModuleAddress, signers[0]) {
			module, _ := k.accountKeeper.AddressCodec().BytesToString(types.ModuleAddress)
			signer, _ := k.accountKeeper.AddressCodec().BytesToString(signers[0])
			return nil, errors.Wrapf(types.ErrInvalidAuthority, "message %d; expected %s, got %s", i, module, signer)
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			return nil, errors.Wrapf(types.ErrInvalidMessage, "no message handler found for %T", msg)
		}
		res, err := handler(sdkCtx, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to execute message %d", i)
		}

		events = append(events, res.Events...)
		results = append(results, res.Data)
	}

	for _, event := range events {
		_ = k.eventService.EventManager(ctx).Emit(ctx, &event)
	}
	return &types.MsgExecuteResponse{Results: results}, nil
}

func (k msgServer) UpdateAuthority(ctx context.Context, msg *types.MsgUpdateAuthority) (*types.MsgUpdateAuthorityResponse, error) {
	authority, err := k.EnsureAuthoritySigner(ctx, msg.Signer)
	if err != nil {
		return nil, err
	}

	newAuthority, err := k.accountKeeper.AddressCodec().StringToBytes(msg.NewAuthority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode new authority address")
	}
	if bytes.Equal(newAuthority, authority) {
		return nil, types.ErrSameAuthority
	}

	err = k.Authority.Set(ctx, newAuthority)
	return &types.MsgUpdateAuthorityResponse{}, err
}

//

func validateMsgs(msgs []sdk.Msg) error {
	for i, raw := range msgs {
		msg, ok := raw.(sdk.HasValidateBasic)
		if !ok {
			continue
		}

		if err := msg.ValidateBasic(); err != nil {
			return errors.Wrapf(err, "msg %d", i)
		}
	}

	return nil
}
