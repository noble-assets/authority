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

package keeper

import (
	"bytes"
	"context"

	"cosmossdk.io/errors"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/authority/types"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Execute(ctx context.Context, msg *types.MsgExecute) (*types.MsgExecuteResponse, error) {
	owner, _ := k.Owner.Get(ctx)
	if msg.Signer != owner {
		return nil, errors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
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
			return nil, errors.Wrapf(types.ErrInvalidOwner, "message %d; expected %s, got %s", i, module, signer)
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

func (k msgServer) TransferOwnership(ctx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	owner, _ := k.Owner.Get(ctx)
	if msg.Signer != owner {
		return nil, errors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	_, err := k.accountKeeper.AddressCodec().StringToBytes(msg.NewOwner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode new owner address")
	}
	if msg.NewOwner == owner {
		return nil, types.ErrSameOwner
	}

	err = k.PendingOwner.Set(ctx, msg.NewOwner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set pending owner in state")
	}

	return &types.MsgTransferOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.OwnershipTransferStarted{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}

func (k msgServer) AcceptOwnership(ctx context.Context, msg *types.MsgAcceptOwnership) (*types.MsgAcceptOwnershipResponse, error) {
	pendingOwner, _ := k.PendingOwner.Get(ctx)
	if pendingOwner == "" {
		return nil, types.ErrNoPendingOwner
	}
	if msg.Signer != pendingOwner {
		return nil, errors.Wrapf(types.ErrInvalidPendingOwner, "expected %s, got %s", pendingOwner, msg.Signer)
	}

	owner, _ := k.Owner.Get(ctx)

	err := k.Owner.Set(ctx, pendingOwner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set owner in state")
	}
	err = k.PendingOwner.Remove(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove pending owner from state")
	}

	return &types.MsgAcceptOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.Signer,
	})
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
