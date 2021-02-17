package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateVote{}, "Themis/CreateVote", nil)

	cdc.RegisterConcrete(&MsgCreatePoll{}, "Themis/CreatePoll", nil)
	cdc.RegisterConcrete(&MsgSetPollDesc{}, "Themis/SetPollDesc", nil)
	cdc.RegisterConcrete(&MsgExtendPollDeadline{}, "Themis/ExtendPollDeadline", nil)

	cdc.RegisterConcrete(&MsgAddAccountVouchers{}, "Themis/AddAccountVouchers", nil)

	cdc.RegisterConcrete(&MsgCreateGroup{}, "Themis/CreateGroup", nil)
	cdc.RegisterConcrete(&MsgSetGroupName{}, "Themis/SetGroupName", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePoll{},
		&MsgSetPollDesc{},
		&MsgExtendPollDeadline{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAccountVouchers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroup{},
		&MsgSetGroupName{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
