package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgSetPollDesc(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPollDesc) (*sdk.Result, error) {
	if !k.PollExists(ctx, msg.ID) { //Checks if group exists
		return nil, sdkerrors.Wrap(types.ErrInvalidPoll, "Invalid poll ID")
	}
	if !msg.Creator.Equals(k.GetGroupOwner(ctx, k.GetPollGroup(ctx, msg.ID))) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetPollDescription(ctx, msg.ID, msg.Description)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
