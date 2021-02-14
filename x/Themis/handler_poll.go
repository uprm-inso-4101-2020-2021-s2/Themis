package Themis

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePoll) (*sdk.Result, error) {

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetGroupOwner(ctx, msg.Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//TODO: charge gas price
	k.CreatePoll(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgSetPollDesc(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSetPollDesc) (*sdk.Result, error) {

	// Checks that the element exists
	if !k.HasPoll(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPollOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//TODO: charge gas price
	k.SetPollDescription(ctx, msg.Id, msg.Description)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgExtendPollDeadline(ctx sdk.Context, k keeper.Keeper, msg *types.MsgExtendPollDeadline) (*sdk.Result, error) {

	// Checks that the element exists
	if !k.HasPoll(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPollOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	if k.GetPoll(ctx, msg.Id).Deadline >= msg.Deadline {
		return nil, sdkerrors.Wrap(types.ErrInvalidPollDate, "Deadlines can only be extended, not shrunk")
	}

	//TODO: charge gas price
	k.ExtendPollDeadline(ctx, msg.Id, msg.Deadline)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
