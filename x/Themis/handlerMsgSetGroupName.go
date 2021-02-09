package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgSetGroupName(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetGroupName) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetGroupOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	//TODO: consume gas upon name change
	k.SetGroupName(ctx, msg.ID, msg.Name)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
