package Themis

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgCreateGroup(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateGroup) (*sdk.Result, error) {
	//TODO: charge gas price

	k.CreateGroup(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgSetGroupName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSetGroupName) (*sdk.Result, error) {
	// Checks that the element exists
	if !k.HasGroup(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetGroupOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//TODO: charge gas price

	k.SetGroupName(ctx, msg.Id, msg.Name)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
