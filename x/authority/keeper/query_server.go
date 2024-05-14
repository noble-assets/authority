package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/authority/x/authority/types"
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
