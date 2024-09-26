package keeper

import (
	"context"

	"cosmossdk.io/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// BeginBlock sends all fees collected in the previous block to the module's owner.
func (k *Keeper) BeginBlock(ctx context.Context) error {
	collector := k.accountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	balance := k.bankKeeper.GetAllBalances(ctx, collector)
	if balance.IsZero() {
		return nil
	}

	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get owner from state")
	}
	address, err := k.accountKeeper.AddressCodec().StringToBytes(owner)
	if err != nil {
		return errors.Wrap(err, "failed to decode owner address")
	}

	return k.bankKeeper.SendCoins(ctx, collector, address, balance)
}
