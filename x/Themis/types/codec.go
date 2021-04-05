package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateVote{}, "Themis/CreateVote", nil)
	cdc.RegisterConcrete(&MsgUpdateVote{}, "Themis/UpdateVote", nil)
	cdc.RegisterConcrete(&MsgDeleteVote{}, "Themis/DeleteVote", nil)

	cdc.RegisterConcrete(&MsgCreatePoll{}, "Themis/CreatePoll", nil)
	cdc.RegisterConcrete(&MsgUpdatePoll{}, "Themis/UpdatePoll", nil)
	cdc.RegisterConcrete(&MsgDeletePoll{}, "Themis/DeletePoll", nil)

	cdc.RegisterConcrete(&MsgCreateGroup{}, "Themis/CreateGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateGroup{}, "Themis/UpdateGroup", nil)
	cdc.RegisterConcrete(&MsgDeleteGroup{}, "Themis/DeleteGroup", nil)

	cdc.RegisterConcrete(&MsgCreateAccount{}, "Themis/CreateAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateAccount{}, "Themis/UpdateAccount", nil)
	cdc.RegisterConcrete(&MsgDeleteAccount{}, "Themis/DeleteAccount", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVote{},
		&MsgUpdateVote{},
		&MsgDeleteVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePoll{},
		&MsgUpdatePoll{},
		&MsgDeletePoll{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroup{},
		&MsgUpdateGroup{},
		&MsgDeleteGroup{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAccount{},
		&MsgUpdateAccount{},
		&MsgDeleteAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
