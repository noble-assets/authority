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

package authority

import (
	"context"
	"encoding/json"
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	modulev1 "github.com/noble-assets/authority/api/module/v1"
	authorityv1 "github.com/noble-assets/authority/api/v1"
	"github.com/noble-assets/authority/client/cli"
	"github.com/noble-assets/authority/keeper"
	"github.com/noble-assets/authority/types"
	"github.com/spf13/cobra"
)

// ConsensusVersion defines the current x/authority module consensus version.
const ConsensusVersion = 1

var (
	_ module.AppModuleBasic      = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ appmodule.HasBeginBlocker  = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasServices         = AppModule{}
)

//

type AppModuleBasic struct{}

func NewAppModuleBasic() AppModuleBasic {
	return AppModuleBasic{}
}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(reg codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var genesis types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genesis); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return genesis.Validate()
}

//

type AppModule struct {
	AppModuleBasic

	keeper        *keeper.Keeper
	accountKeeper types.AccountKeeper
}

func NewAppModule(keeper *keeper.Keeper, accountKeeper types.AccountKeeper) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
	}
}

func (AppModule) IsOnePerModuleType() {}

func (AppModule) IsAppModule() {}

func (m AppModule) BeginBlock(ctx context.Context) error {
	return m.keeper.BeginBlock(ctx)
}

func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

func (m AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
	var genesis types.GenesisState
	cdc.MustUnmarshalJSON(bz, &genesis)

	InitGenesis(ctx, m.keeper, m.accountKeeper, genesis)
}

func (m AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesis := ExportGenesis(ctx, m.keeper)
	return cdc.MustMarshalJSON(genesis)
}

func (m AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(m.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(m.keeper))
}

//

func (AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: authorityv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				// NOTE: Execute is handled via a custom CLI command.
				{
					RpcMethod:      "TransferOwnership",
					Use:            "transfer-ownership [new-owner]",
					Short:          "Transfer ownership of the module",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "new_owner"}},
				},
				{
					RpcMethod: "AcceptOwnership",
					Use:       "accept-ownership",
					Short:     "Accept ownership of the module",
				},
			},
			EnhanceCustomCommand: true,
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: authorityv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Owner",
					Use:       "owner",
					Short:     "Query owner of the module",
				},
				{
					RpcMethod: "PendingOwner",
					Use:       "pending-owner",
					Short:     "Query pending owner of the module",
				},
			},
		},
	}
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

//

func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config       *modulev1.Module
	Cdc          codec.Codec
	StoreService store.KVStoreService
	EventService event.Service

	Router        baseapp.MessageRouter
	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
}

type ModuleOutputs struct {
	depinject.Out

	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.EventService,
		in.Router,
		in.AccountKeeper,
		in.BankKeeper,
	)
	m := NewAppModule(k, in.AccountKeeper)

	return ModuleOutputs{Keeper: k, Module: m}
}
