package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExecute{}, "noble/Execute", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "noble/TransferOwnership", nil)
	cdc.RegisterConcrete(&MsgAcceptOwnership{}, "noble/AcceptOwnership", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecute{},
		&MsgTransferOwnership{},
		&MsgAcceptOwnership{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
