package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVoucher{}

type MsgCreateVoucher struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Group   string         `json:"group" yaml:"group"`
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
}

func NewMsgCreateVoucher(creator sdk.AccAddress, owner sdk.AccAddress, group string) MsgCreateVoucher {
	return MsgCreateVoucher{
		Creator: creator,
		Group:   group,
		Owner:   owner,
	}
}

func (msg MsgCreateVoucher) Route() string {
	return RouterKey
}

func (msg MsgCreateVoucher) Type() string {
	return "CreateVoucher"
}

func (msg MsgCreateVoucher) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateVoucher) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateVoucher) ValidateBasic() error {
	if msg.Creator.Empty() || msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Group == "" {
		return sdkerrors.Wrap(ErrInvalidGroup, "Group can't be empty")
	}
	return nil
}
