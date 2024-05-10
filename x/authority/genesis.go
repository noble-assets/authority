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

	if genesis.PendingAddress != "" {
		pendingAuthority, err := accountKeeper.AddressCodec().StringToBytes(genesis.PendingAddress)
		if err != nil {
			panic("failed to decode pending authority address")
		}
		err = k.PendingAuthority.Set(ctx, pendingAuthority)
		if err != nil {
			panic("failed to set pending authority in state")
		}
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

	pendingAddress := ""
	pendingAuthority, err := k.PendingAuthority.Get(ctx)
	if err == nil {
		pendingAddress, err = accountKeeper.AddressCodec().BytesToString(pendingAuthority)
		if err != nil {
			panic("failed to encode pending authority address")
		}
	}

	return &types.GenesisState{Address: address, PendingAddress: pendingAddress}
}
