package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreatePoll{}, "Themis/CreatePoll", nil)
	cdc.RegisterConcrete(MsgExtendPollDeadline{}, "Themis/ExtendPollDeadline", nil)
	cdc.RegisterConcrete(MsgSetPollDesc{}, "Themis/SetPollDesc", nil)
	cdc.RegisterConcrete(MsgAccountAddVotes{}, "Themis/CreateVoucher", nil)
	cdc.RegisterConcrete(MsgCreateGroup{}, "Themis/CreateGroup", nil)
	cdc.RegisterConcrete(MsgSetGroupName{}, "Themis/SetGroupName", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
