package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGroup{}

type MsgCreateGroup struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
}

func NewMsgCreateGroup(creator sdk.AccAddress, name string) MsgCreateGroup {
	return MsgCreateGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg MsgCreateGroup) Route() string {
	return RouterKey
}

func (msg MsgCreateGroup) Type() string {
	return "CreateGroup"
}

func (msg MsgCreateGroup) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGroup) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(ErrNameEmpty, "Name can't be empty")
	}
	return nil
}
