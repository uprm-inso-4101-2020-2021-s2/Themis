package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgExtendPollDeadline{}

type MsgExtendPollDeadline struct {
	ID       string         `json:"id" yaml:"id"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Deadline int            `json:"deadline" yaml:"deadline"`
}

func NewMsgExtendPollDeadline(creator sdk.AccAddress, id string, deadline int) MsgExtendPollDeadline {
	return MsgExtendPollDeadline{
		ID:       id,
		Creator:  creator,
		Deadline: deadline,
	}
}

func (msg MsgExtendPollDeadline) Route() string {
	return RouterKey
}

func (msg MsgExtendPollDeadline) Type() string {
	return "ExtendPollDeadline"
}

func (msg MsgExtendPollDeadline) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgExtendPollDeadline) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgExtendPollDeadline) ValidateBasic() error {
	if msg.ID == "" {
		return sdkerrors.Wrap(ErrInvalidPoll, "poll id can't be empty")
	}
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	//TODO: check deadline already reached here
	return nil
}
