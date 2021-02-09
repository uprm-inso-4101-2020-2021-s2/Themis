package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetGroupName{}

type MsgSetGroupName struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
}

func NewMsgSetGroupName(creator sdk.AccAddress, id string, name string) MsgSetGroupName {
	return MsgSetGroupName{
		ID:      id,
		Creator: creator,
		Name:    name,
	}
}

func (msg MsgSetGroupName) Route() string {
	return RouterKey
}

func (msg MsgSetGroupName) Type() string {
	return "SetGroup"
}

func (msg MsgSetGroupName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetGroupName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetGroupName) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(ErrNameEmpty, "name can't be empty")
	}
	return nil
}
