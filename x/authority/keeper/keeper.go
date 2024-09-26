package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/store"
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

	Schema       collections.Schema
	Owner        collections.Item[string]
	PendingOwner collections.Item[string]

	router        baseapp.MessageRouter
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewKeeper(
	cdc codec.Codec,
	logger log.Logger,
	storeService store.KVStoreService,
	eventService event.Service,
	router baseapp.MessageRouter,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	builder := collections.NewSchemaBuilder(storeService)

	keeper := &Keeper{
		cdc:          cdc,
		logger:       logger,
		storeService: storeService,
		eventService: eventService,

		Owner:        collections.NewItem(builder, types.OwnerKey, "owner", collections.StringValue),
		PendingOwner: collections.NewItem(builder, types.PendingOwnerKey, "pending_owner", collections.StringValue),

		router:        router,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}

	schema, err := builder.Build()
	if err != nil {
		panic(err)
	}

	keeper.Schema = schema
	return keeper
}
