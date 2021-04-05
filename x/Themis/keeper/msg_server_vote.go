package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func (k msgServer) CreateVote(goCtx context.Context, msg *types.MsgCreateVote) (*types.MsgCreateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check that user hasnt already voted
	if k.UserVoted(ctx, msg.Creator, msg.Poll) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Already voted")
	}

	// Check that user has permission to vote
	if !k.AccountInGroup(ctx, k.GetAccountId(ctx, msg.Creator), k.GetPoll(ctx, msg.Poll).Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "User unauthorized to vote in this poll")
	}

	// Check option exists
	if !k.PollVoteExists(ctx, msg.Poll, msg.Option) {
		return nil, sdkerrors.Wrap(types.ErrInvalidVotes, "Option doesnt exist")
	}

	id := k.AppendVote(
		ctx,
		msg.Creator,
		msg.Poll,
		msg.Option,
	)

	k.AddPollVote(ctx, msg.Poll, msg.Option)

	return &types.MsgCreateVoteResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateVote(goCtx context.Context, msg *types.MsgUpdateVote) (*types.MsgUpdateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	if !k.HasVote(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetVoteOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	vote := k.GetVote(ctx, msg.Id)

	// Check option exists
	if !k.PollVoteExists(ctx, vote.Poll, msg.Option) {
		return nil, sdkerrors.Wrap(types.ErrInvalidVotes, "Option doesnt exist")
	}

	k.RemovePollVote(ctx, vote.Poll, vote.Option)
	k.AddPollVote(ctx, vote.Poll, msg.Option)

	vote.Option = msg.Option

	k.SetVote(ctx, vote)

	return &types.MsgUpdateVoteResponse{}, nil
}

func (k msgServer) DeleteVote(goCtx context.Context, msg *types.MsgDeleteVote) (*types.MsgDeleteVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasVote(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetVoteOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePollVote(ctx, k.GetVote(ctx, msg.Id).Poll, k.GetVote(ctx, msg.Id).Option)
	k.RemoveVote(ctx, msg.Id)

	return &types.MsgDeleteVoteResponse{}, nil
}
