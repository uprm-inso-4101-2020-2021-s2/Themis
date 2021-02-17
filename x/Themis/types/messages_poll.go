package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/types/time"
)

var maxTitleSize = 60
var maxDescSize = 140

var _ sdk.Msg = &MsgCreatePoll{}

func NewMsgCreatePoll(creator string, group string, title string, description string, options []string, deadline int64) *MsgCreatePoll {
	return &MsgCreatePoll{
		Creator:     creator,
		Group:       group,
		Title:       title,
		Description: description,
		Options:     options,
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
	if msg.Title == "" || len(msg.Title) > maxTitleSize {
		return sdkerrors.Wrapf(ErrInvalidPollTitle, "Poll title can't be empty or longer than (%s) characters", maxTitleSize)
	}
	if msg.Description == "" || len(msg.Description) > maxDescSize {
		return sdkerrors.Wrapf(ErrInvalidPollTitle, "Poll description can't be empty or longer than (%s) characters", maxDescSize)
	}
	if len(msg.Options) <= 0 {
		return sdkerrors.Wrapf(ErrInvalidPollOptns, "Options can't empty")
	}
	if msg.Deadline <= time.Now().Unix() {
		return sdkerrors.Wrapf(ErrInvalidPollDate, "Deadline must be a date in the future")
	}
	return nil
}

var _ sdk.Msg = &MsgSetPollDesc{}

func NewMsgSetPollDesc(creator string, id string, description string) *MsgSetPollDesc {
	return &MsgSetPollDesc{
		Creator:     creator,
		Id:          id,
		Description: description,
	}
}

func (msg *MsgSetPollDesc) Route() string {
	return RouterKey
}

func (msg *MsgSetPollDesc) Type() string {
	return "SetPollDesc"
}

func (msg *MsgSetPollDesc) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPollDesc) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPollDesc) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Id == "" {
		return sdkerrors.Wrapf(ErrInvalidPoll, "Poll can't empty")
	}
	if msg.Description == "" || len(msg.Description) > maxDescSize {
		return sdkerrors.Wrapf(ErrInvalidPollDesc, "Poll description can't be empty or longer than (%s) characters", maxDescSize)
	}
	return nil
}

var _ sdk.Msg = &MsgExtendPollDeadline{}

func NewMsgExtendPollDeadline(creator string, id string, deadline int64) *MsgExtendPollDeadline {
	return &MsgExtendPollDeadline{
		Creator:  creator,
		Id:       id,
		Deadline: deadline,
	}
}

func (msg *MsgExtendPollDeadline) Route() string {
	return RouterKey
}

func (msg *MsgExtendPollDeadline) Type() string {
	return "ExtendPollDeadline"
}

func (msg *MsgExtendPollDeadline) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExtendPollDeadline) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExtendPollDeadline) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Id == "" {
		return sdkerrors.Wrapf(ErrInvalidPoll, "Poll can't empty")
	}
	if msg.Deadline <= time.Now().Unix() {
		return sdkerrors.Wrapf(ErrInvalidPollDate, "Deadline must be a date in the future")
	}
	return nil
}
