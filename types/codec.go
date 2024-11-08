// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

import (
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExecute{}, "noble/Execute", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "noble/TransferOwnership", nil)
	cdc.RegisterConcrete(&MsgAcceptOwnership{}, "noble/AcceptOwnership", nil)
	cdc.RegisterConcrete(&upgradetypes.MsgSoftwareUpgrade{}, "noble/SoftwareUpgrade", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecute{},
		&MsgTransferOwnership{},
		&MsgAcceptOwnership{},
		&upgradetypes.MsgSoftwareUpgrade{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
