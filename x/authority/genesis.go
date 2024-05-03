package authority

import (
	"context"

	"github.com/noble-assets/authority/x/authority/keeper"
	"github.com/noble-assets/authority/x/authority/types"
)

func InitGenesis(ctx context.Context, k *keeper.Keeper, accountKeeper types.AccountKeeper, genesis types.GenesisState) {
	authority, err := accountKeeper.AddressCodec().StringToBytes(genesis.Address)
	if err != nil {
		panic("failed to decode authority address")
	}
	err = k.Authority.Set(ctx, authority)
	if err != nil {
		panic("failed to set authority in state")
	}
}

func ExportGenesis(ctx context.Context, k *keeper.Keeper, accountKeeper types.AccountKeeper) *types.GenesisState {
	authority, err := k.Authority.Get(ctx)
	if err != nil {
		panic("failed to get authority from state")
	}
	address, err := accountKeeper.AddressCodec().BytesToString(authority)
	if err != nil {
		panic("failed to encode authority address")
	}

	return &types.GenesisState{Address: address}
}
