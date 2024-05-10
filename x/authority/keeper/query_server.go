package keeper

import (
	"context"

	"cosmossdk.io/errors"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/authority/x/authority/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	*Keeper
}

func NewQueryServer(keeper *Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

func (k queryServer) Address(ctx context.Context, req *types.QueryAddress) (*types.QueryAddressResponse, error) {
	if req == nil {
		return nil, errorstypes.ErrInvalidRequest
	}

	authority, err := k.Authority.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve authority from state")
	}
	address, err := k.accountKeeper.AddressCodec().BytesToString(authority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode authority address")
	}

	pendingAddress := ""
	pendingAuthority, err := k.PendingAuthority.Get(ctx)
	if err == nil {
		pendingAddress, err = k.accountKeeper.AddressCodec().BytesToString(pendingAuthority)
		if err != nil {
			return nil, errors.Wrap(err, "failed to encode pending authority address")
		}
	}

	return &types.QueryAddressResponse{Address: address, PendingAddress: pendingAddress}, nil
}
