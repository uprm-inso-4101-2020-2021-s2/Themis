package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAccountAddVotes{}

type MsgAccountAddVotes struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Group   string         `json:"group" yaml:"group"`
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Amount  int            `json:"amount" yaml:"amount"`
}

func NewMsgAccountAddVotes(creator sdk.AccAddress, owner sdk.AccAddress, group string, votes int) MsgAccountAddVotes {
	return MsgAccountAddVotes{
		Creator: creator,
		Group:   group,
		Owner:   owner,
		Amount:  votes,
	}
}

func (msg MsgAccountAddVotes) Route() string {
	return RouterKey
}

func (msg MsgAccountAddVotes) Type() string {
	return "AddVotes"
}

func (msg MsgAccountAddVotes) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgAccountAddVotes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgAccountAddVotes) ValidateBasic() error {
	if msg.Creator.Empty() || msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Group == "" {
		return sdkerrors.Wrap(ErrInvalidGroup, "Group can't be empty")
	}
	if msg.Amount < 0 {
		return sdkerrors.Wrap(ErrVoteNumberInvalid, "Vote amount can't be invalid")
	}
	return nil
}
