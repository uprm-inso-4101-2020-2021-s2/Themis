package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVote{}

func NewMsgCreateVote(creator string, poll string, option int32) *MsgCreateVote {
	return &MsgCreateVote{
		Creator: creator,
		Poll:    poll,
		Option:  option,
	}
}

func (msg *MsgCreateVote) Route() string {
	return RouterKey
}

func (msg *MsgCreateVote) Type() string {
	return "CreateVote"
}

func (msg *MsgCreateVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Poll == "" {
		return sdkerrors.Wrapf(ErrInvalidPoll, "Poll can't empty")
	}
	if msg.Option < 0 {
		return sdkerrors.Wrapf(ErrInvalidVoteOption, "vote cannot be less than 0")
	}
	return nil
}
