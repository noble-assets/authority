// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package mocks

import (
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/noble-assets/authority/keeper"
	"github.com/noble-assets/authority/types"
)

func AuthorityKeeper() (*keeper.Keeper, sdk.Context) {
	return AuthorityKeeperWithBank(BankKeeper{})
}

func AuthorityKeeperWithBank(bank types.BankKeeper) (*keeper.Keeper, sdk.Context) {
	keys := storetypes.NewKVStoreKeys(
		types.ModuleName,
		upgradetypes.StoreKey,
	)

	// key := storetypes.NewKVStoreKey(types.ModuleName)
	tkey := storetypes.NewTransientStoreKey("transient_authority")

	cfg := MakeTestEncodingConfig("noble", upgrade.AppModuleBasic{})
	router := baseapp.NewMsgServiceRouter()
	router.SetInterfaceRegistry(cfg.InterfaceRegistry)
	upgradetypes.RegisterMsgServer(router, UpgradeKeeper{})

	return keeper.NewKeeper(
		cfg.Codec,
		runtime.NewKVStoreService(keys[types.ModuleName]),
		runtime.ProvideEventService(),
		router,
		AccountKeeper{},
		bank,
		upgradekeeper.Keeper{},
	), testutil.DefaultContext(keys[types.ModuleName], tkey)
}

// MakeTestEncodingConfig is a modified testutil.MakeTestEncodingConfig that
// sets a custom Bech32 prefix in the interface registry.
func MakeTestEncodingConfig(prefix string, modules ...module.AppModuleBasic) moduletestutil.TestEncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectestutil.CodecOptions{
		AccAddressPrefix: prefix,
	}.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	encCfg := moduletestutil.TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             aminoCodec,
	}

	mb := module.NewBasicManager(modules...)

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	mb.RegisterLegacyAminoCodec(encCfg.Amino)
	mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}
