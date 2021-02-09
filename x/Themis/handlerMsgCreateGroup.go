package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgCreateGroup(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGroup) (*sdk.Result, error) {
	// TODO: instead of removing money from funds, just remove money from the group's creator
	k.CreateGroup(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
