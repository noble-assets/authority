package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExecute{}, "noble/authority/Execute", nil)
	cdc.RegisterConcrete(&MsgUpdateAuthority{}, "noble/authority/UpdateAuthority", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgExecute{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateAuthority{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
