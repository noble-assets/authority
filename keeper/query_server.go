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
	"context"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/authority/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	*Keeper
}

func NewQueryServer(keeper *Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

func (k queryServer) Owner(ctx context.Context, req *types.QueryOwner) (*types.QueryOwnerResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	owner, err := k.Keeper.Owner.Get(ctx)

	return &types.QueryOwnerResponse{Owner: owner}, err
}

func (k queryServer) PendingOwner(ctx context.Context, req *types.QueryPendingOwner) (*types.QueryPendingOwnerResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	pendingOwner, _ := k.Keeper.PendingOwner.Get(ctx)

	return &types.QueryPendingOwnerResponse{PendingOwner: pendingOwner}, nil
}
