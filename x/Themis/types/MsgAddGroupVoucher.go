package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddGroupVoucher{}

type MsgAddGroupVoucher struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgAddGroupVoucher(creator sdk.AccAddress, id string) MsgAddGroupVoucher {
	return MsgAddGroupVoucher{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgAddGroupVoucher) Route() string {
	return RouterKey
}

func (msg MsgAddGroupVoucher) Type() string {
	return "AddGroupVouchers"
}

func (msg MsgAddGroupVoucher) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgAddGroupVoucher) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgAddGroupVoucher) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
