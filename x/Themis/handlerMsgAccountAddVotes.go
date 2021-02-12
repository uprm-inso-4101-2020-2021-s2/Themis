package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgAccountAddVotes(ctx sdk.Context, k keeper.Keeper, msg types.MsgAccountAddVotes) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetGroupOwner(ctx, msg.Group)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	if k.UserGroupAccountExists(ctx, msg.Owner.String(), string(msg.Group)) {
		k.AddToAccount(ctx, msg.Owner.String(), string(msg.Group), msg.Amount)
	} else {
		k.CreateAccount(ctx, msg)
		k.AddGroupVoucher(ctx, msg.Group) // Add one more voucher
	}
	// TODO: gas price to creator

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
