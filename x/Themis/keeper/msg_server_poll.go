package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func (k msgServer) CreatePoll(goCtx context.Context, msg *types.MsgCreatePoll) (*types.MsgCreatePollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendPoll(
		ctx,
		msg.Creator,
		msg.Name,
		msg.Group,
		msg.Votes,
		msg.Description,
		msg.Deadline,
	)

	// Checks that the group element exists
	if !k.HasGroup(ctx, msg.Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Group key %d doesn't exist", msg.Group))
	}

	// Check that creator is allowed to post a poll
	if msg.Creator != k.GetGroupOwner(ctx, msg.Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//TODO: check that the deadline hasn't passed

	return &types.MsgCreatePollResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePoll(goCtx context.Context, msg *types.MsgUpdatePoll) (*types.MsgUpdatePollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	if !k.HasPoll(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPollOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	poll := k.GetPoll(ctx, msg.Id)
	poll.Description = msg.Description

	if poll.Deadline > msg.Deadline {
		return nil, sdkerrors.Wrap(types.ErrInvalidDate, fmt.Sprintf("Date %d is sooner than original posted date, deadlines can only be extended", msg.Deadline))
	}
	//TODO: check that the deadline hasn't passed
	poll.Deadline = msg.Deadline

	k.SetPoll(ctx, poll)

	return &types.MsgUpdatePollResponse{}, nil
}

func (k msgServer) DeletePoll(goCtx context.Context, msg *types.MsgDeletePoll) (*types.MsgDeletePollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasPoll(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetPollOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePoll(ctx, msg.Id)

	return &types.MsgDeletePollResponse{}, nil
}
