// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/noble-assets/authority/types"
)

type Keeper struct {
	cdc          codec.Codec
	storeService store.KVStoreService
	eventService event.Service

	Schema       collections.Schema
	Owner        collections.Item[string]
	PendingOwner collections.Item[string]

	router        baseapp.MessageRouter
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	upgradeKeeper types.UpgradeKeeper
}

func NewKeeper(
	cdc codec.Codec,
	storeService store.KVStoreService,
	eventService event.Service,
	router baseapp.MessageRouter,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	upgradeKeeper types.UpgradeKeeper,
) *Keeper {
	builder := collections.NewSchemaBuilder(storeService)

	keeper := &Keeper{
		cdc:          cdc,
		storeService: storeService,
		eventService: eventService,

		Owner:        collections.NewItem(builder, types.OwnerKey, "owner", collections.StringValue),
		PendingOwner: collections.NewItem(builder, types.PendingOwnerKey, "pending_owner", collections.StringValue),

		router:        router,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		upgradeKeeper: upgradeKeeper,
	}

	schema, err := builder.Build()
	if err != nil {
		panic(err)
	}

	keeper.Schema = schema
	return keeper
}
