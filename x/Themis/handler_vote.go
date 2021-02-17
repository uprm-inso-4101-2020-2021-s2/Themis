package Themis

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/types/time"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func handleMsgCreateVote(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateVote) (*sdk.Result, error) {
	// Check if poll exists
	if !k.HasPoll(ctx, msg.Poll) {
		return nil, sdkerrors.Wrap(types.ErrInvalidPoll, fmt.Sprintf("Poll #{msg.Poll} doesn't exist"))
	}
	if int32(len(k.GetPollOptions(ctx, msg.Poll))) < msg.Option {
		return nil, sdkerrors.Wrap(types.ErrInvalidVoteOption, fmt.Sprintf("Vote option exceeded total options"))
	}
	if k.GetPollDeadline(ctx, msg.Poll) < time.Now().Unix() {
		return nil, sdkerrors.Wrap(types.ErrPollDateReacher, fmt.Sprintf("Poll deadline already reached"))
	}
	group := k.GetPollGroup(ctx, msg.Poll)
	if !k.AccountExistsInGroup(ctx, msg.Creator, group) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccount, fmt.Sprintf("Account #{msg.Creator} isn't in the poll's group"))
	}
	acc := k.GetAccountId(ctx, msg.Creator, group)
	if k.GetAccountVouchers(ctx, acc) <= 0 {
		return nil, sdkerrors.Wrap(types.ErrNoAccountFunds, fmt.Sprintf("Account #{msg.Creator} doesn't have enough vouchers"))
	}
	k.EditAccountVouchers(ctx, acc, -1)

	//TODO: charge gas price
	k.CreateVote(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
