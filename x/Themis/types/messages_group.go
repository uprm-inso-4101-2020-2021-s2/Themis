package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGroup{}

//TODO: group IDs do not reflect the creator, update this so the creator is part of the ID, this allows for user based queries in groups

func NewMsgCreateGroup(creator string, name string) *MsgCreateGroup {
	return &MsgCreateGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgCreateGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroup) Type() string {
	return "CreateGroup"
}

func (msg *MsgCreateGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrNameEmpty, "name can't be empty")
	}
	return nil
}

var _ sdk.Msg = &MsgSetGroupName{}

func NewMsgSetGroupName(creator string, id string, name string) *MsgSetGroupName {
	return &MsgSetGroupName{
		Id:      id,
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgSetGroupName) Route() string {
	return RouterKey
}

func (msg *MsgSetGroupName) Type() string {
	return "SetGroupName"
}

func (msg *MsgSetGroupName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetGroupName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetGroupName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrNameEmpty, "name can't be empty")
	}
	if msg.Id == "" {
		return sdkerrors.Wrapf(ErrInvalidGroup, "group ID can't be empty")
	}
	return nil
}
