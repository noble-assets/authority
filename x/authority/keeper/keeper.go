package keeper

import (
	"bytes"
	"context"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/store"
	"cosmossdk.io/errors"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/noble-assets/authority/x/authority/types"
)

type Keeper struct {
	cdc          codec.Codec
	logger       log.Logger
	storeService store.KVStoreService
	eventService event.Service

	Schema           collections.Schema
	Authority        collections.Item[[]byte]
	PendingAuthority collections.Item[[]byte]

	router        baseapp.MessageRouter
	accountKeeper types.AccountKeeper
}

func NewKeeper(
	cdc codec.Codec,
	logger log.Logger,
	storeService store.KVStoreService,
	eventService event.Service,
	router baseapp.MessageRouter,
	accountKeeper types.AccountKeeper,
) *Keeper {
	builder := collections.NewSchemaBuilder(storeService)

	keeper := &Keeper{
		cdc:          cdc,
		logger:       logger,
		storeService: storeService,
		eventService: eventService,

		Authority:        collections.NewItem(builder, types.AuthorityKey, "authority", collections.BytesValue),
		PendingAuthority: collections.NewItem(builder, types.PendingAuthorityKey, "pending_authority", collections.BytesValue),

		router:        router,
		accountKeeper: accountKeeper,
	}

	schema, err := builder.Build()
	if err != nil {
		panic(err)
	}

	keeper.Schema = schema
	return keeper
}

// EnsureAuthoritySigner is helper function that checks that the signer of a
// message is the current authority.
func (k *Keeper) EnsureAuthoritySigner(ctx context.Context, signerAddress string) ([]byte, error) {
	authority, err := k.Authority.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve authority from state")
	}
	signer, err := k.accountKeeper.AddressCodec().StringToBytes(signerAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode signer address")
	}
	if !bytes.Equal(signer, authority) {
		address, _ := k.accountKeeper.AddressCodec().BytesToString(authority)
		return nil, errors.Wrapf(types.ErrInvalidAuthority, "expected %s, got %s", address, signerAddress)
	}

	return authority, nil
}
