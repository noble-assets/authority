package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExecute{}, "noble/Execute", nil)
	cdc.RegisterConcrete(&MsgTransferAuthority{}, "noble/TransferAuthority", nil)
	cdc.RegisterConcrete(&MsgAcceptAuthority{}, "noble/AcceptAuthority", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecute{},
		&MsgTransferAuthority{},
		&MsgAcceptAuthority{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
