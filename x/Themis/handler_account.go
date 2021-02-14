package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgAddAccountVouchers(ctx sdk.Context, k keeper.Keeper, msg *types.MsgAddAccountVouchers) (*sdk.Result, error) {

	// Checks if the the msg sender is the same as the current owner
	if msg.GroupOwner != k.GetGroupOwner(ctx, msg.Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	//TODO: pay gas from group owner account

	// Checks that the element exists
	if k.AccountExistsInGroup(ctx, msg.User, msg.Group) {
		k.SetAccountVoucher(ctx, k.NewKey(msg.User, msg.Group), msg.Vouchers)
	} else {
		k.CreateAccount(ctx, *msg)
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
