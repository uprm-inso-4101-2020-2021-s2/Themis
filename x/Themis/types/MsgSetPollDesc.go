package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPollDesc{}

type MsgSetPollDesc struct {
	ID          string         `json:"id" yaml:"id"`
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	Description string         `json:"description" yaml:"description"`
}

func NewMsgSetPollDesc(creator sdk.AccAddress, id string, desc string) MsgSetPollDesc {
	return MsgSetPollDesc{
		ID:          id,
		Creator:     creator,
		Description: desc,
	}
}

func (msg MsgSetPollDesc) Route() string {
	return RouterKey
}

func (msg MsgSetPollDesc) Type() string {
	return "SetPollDesc"
}

func (msg MsgSetPollDesc) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetPollDesc) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetPollDesc) ValidateBasic() error {
	if msg.ID == "" {
		return sdkerrors.Wrap(ErrInvalidPoll, "poll id can't be empty")
	}
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if len(msg.Description) > 140 {
		return sdkerrors.Wrap(ErrDescLimitPassed, "Description must be 140 characters or less")
	}
	return nil
}
