package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func (k msgServer) CreateGroup(goCtx context.Context, msg *types.MsgCreateGroup) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendGroup(
		ctx,
		msg.Name,
		msg.Owner,
	)

	return &types.MsgCreateGroupResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateGroup(goCtx context.Context, msg *types.MsgUpdateGroup) (*types.MsgUpdateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var group = types.Group{
		Id:    msg.Id,
		Name:  msg.Name,
		Owner: msg.NewOwner,
	}

	// Checks that the element exists
	if !k.HasGroup(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Owner != k.GetGroupOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetGroup(ctx, group)

	return &types.MsgUpdateGroupResponse{}, nil
}

func (k msgServer) DeleteGroup(goCtx context.Context, msg *types.MsgDeleteGroup) (*types.MsgDeleteGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasGroup(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetGroupOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveGroup(ctx, msg.Id)

	return &types.MsgDeleteGroupResponse{}, nil
}
