package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoll{}

type MsgCreatePoll struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	Group       string         `json:"group" yaml:"group"`
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Options     []string       `json:"options" yaml:"options"`
	Deadline    int            `json:"deadline" yaml:"deadline"`
}

func NewMsgCreatePoll(creator sdk.AccAddress, group string, title string, description string, options []string, deadline int) MsgCreatePoll {
	return MsgCreatePoll{
		Creator:     creator,
		Group:       group,
		Title:       title,
		Description: description,
		Options:     options,
		Deadline:    deadline,
	}
}

func (msg MsgCreatePoll) Route() string {
	return RouterKey
}

func (msg MsgCreatePoll) Type() string {
	return "CreatePoll"
}

func (msg MsgCreatePoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePoll) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Title == "" {
		return sdkerrors.Wrap(ErrTitleEmpty, "Title can't be empty")
	}
	if len(msg.Options) <= 0 {
		return sdkerrors.Wrap(ErrOptionsEmpty, "Poll must have at least one option")
	}
	if len(msg.Description) > 140 {
		return sdkerrors.Wrap(ErrDescLimitPassed, "Description must be 140 characters or less")
	}
	//TODO: check deadline already reached here
	return nil
}
