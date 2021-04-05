package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoll{}

func NewMsgCreatePoll(creator string, name string, group uint64, votes []string, description string, deadline uint64) *MsgCreatePoll {
	return &MsgCreatePoll{
		Creator:     creator,
		Name:        name,
		Group:       group,
		Votes:       votes,
		Description: description,
		Deadline:    deadline,
	}
}

func (msg *MsgCreatePoll) Route() string {
	return RouterKey
}

func (msg *MsgCreatePoll) Type() string {
	return "CreatePoll"
}

func (msg *MsgCreatePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" || len(msg.Name) > MaxNameSize {
		return sdkerrors.Wrapf(ErrInvalidName, "Name either too small or too large (%s)", msg.Name)
	}
	if len(msg.Description) > MaxDescSize {
		return sdkerrors.Wrapf(ErrInvalidName, "Description too large (%s)", msg.Description)
	}
	if len(msg.Votes) > MaxOptions || len(msg.Votes) <= 1 {
		return sdkerrors.Wrapf(ErrInvalidVotes, "Votes either too few or too many (%s)", msg.Votes)
	}
	if msg.Deadline <= 0 {
		return sdkerrors.Wrapf(ErrInvalidDate, "Date must be UTC time (%s)", msg.Deadline)
	}
	//TODO: check that no identical vote strings are not passed in array, check that vote strings donot pass a size
	return nil
}

var _ sdk.Msg = &MsgUpdatePoll{}

func NewMsgUpdatePoll(creator string, id uint64, description string, deadline uint64) *MsgUpdatePoll {
	return &MsgUpdatePoll{
		Id:          id,
		Creator:     creator,
		Description: description,
		Deadline:    deadline,
	}
}

func (msg *MsgUpdatePoll) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePoll) Type() string {
	return "UpdatePoll"
}

func (msg *MsgUpdatePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Deadline <= 0 {
		return sdkerrors.Wrapf(ErrInvalidDate, "Date must be UTC time (%s)", msg.Deadline)
	}
	if len(msg.Description) > MaxDescSize {
		return sdkerrors.Wrapf(ErrInvalidName, "Description too large (%s)", msg.Description)
	}
	return nil
}

var _ sdk.Msg = &MsgCreatePoll{}

func NewMsgDeletePoll(creator string, id uint64) *MsgDeletePoll {
	return &MsgDeletePoll{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePoll) Route() string {
	return RouterKey
}

func (msg *MsgDeletePoll) Type() string {
	return "DeletePoll"
}

func (msg *MsgDeletePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
