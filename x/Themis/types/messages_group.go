package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGroup{}

func NewMsgCreateGroup(name string, owner string) *MsgCreateGroup {
	return &MsgCreateGroup{
		Name:  name,
		Owner: owner,
	}
}

func (msg *MsgCreateGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroup) Type() string {
	return "CreateGroup"
}

func (msg *MsgCreateGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
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
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" || len(msg.Name) > MaxNameSize {
		return sdkerrors.Wrapf(ErrInvalidName, "Name either too small or too large (%s)", msg.Name)
	}
	return nil
}

var _ sdk.Msg = &MsgInviteToGroup{}

func NewMsgInviteToGroup(group uint64, invited uint64, owner string) *MsgInviteToGroup {
	return &MsgInviteToGroup{
		Group:   group,
		Invited: invited,
		Owner:   owner,
	}
}

func (msg *MsgInviteToGroup) Route() string {
	return RouterKey
}

func (msg *MsgInviteToGroup) Type() string {
	return "InviteToGroup"
}

func (msg *MsgInviteToGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInviteToGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInviteToGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGroup{}

func NewMsgUpdateGroup(id uint64, name string, owner string, newOwner string) *MsgUpdateGroup {
	return &MsgUpdateGroup{
		Id:       id,
		Name:     name,
		Owner:    owner,
		NewOwner: newOwner,
	}
}

func (msg *MsgUpdateGroup) Route() string {
	return RouterKey
}

func (msg *MsgUpdateGroup) Type() string {
	return "UpdateGroup"
}

func (msg *MsgUpdateGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" || len(msg.Name) > MaxNameSize {
		return sdkerrors.Wrapf(ErrInvalidName, "Name either too small or too large (%s)", msg.Name)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateGroup{}

func NewMsgDeleteGroup(creator string, id uint64) *MsgDeleteGroup {
	return &MsgDeleteGroup{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteGroup) Route() string {
	return RouterKey
}

func (msg *MsgDeleteGroup) Type() string {
	return "DeleteGroup"
}

func (msg *MsgDeleteGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
