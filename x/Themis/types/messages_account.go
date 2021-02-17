package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddAccountVouchers{}

func NewMsgAddAccountVouchers(groupOwner string, user string, group string, vouchers int64) *MsgAddAccountVouchers {
	return &MsgAddAccountVouchers{
		GroupOwner: groupOwner,
		User:       user,
		Group:      group,
		Vouchers:   vouchers,
	}
}

func (msg *MsgAddAccountVouchers) Route() string {
	return RouterKey
}

func (msg *MsgAddAccountVouchers) Type() string {
	return "AddAccountVouchers"
}

func (msg *MsgAddAccountVouchers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.GroupOwner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAccountVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAccountVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GroupOwner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	if msg.Group == "" {
		return sdkerrors.Wrapf(ErrInvalidGroup, "Group can't be empty")
	}
	if msg.Vouchers < 0 {
		return sdkerrors.Wrapf(ErrVoteNumberInvalid, "Vote amount can't be invalid")
	}

	return nil
}
