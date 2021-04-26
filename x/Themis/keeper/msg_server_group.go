package keeper

import (
	"context"
	"fmt"
	"github.com/tendermint/tendermint/types/time"

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

	var acc = k.GetAccountFromAddr(ctx, msg.Owner)

	k.AddAccountGroup(ctx, acc.Id, id, uint64(time.Now().UTC().Unix()))

	return &types.MsgCreateGroupResponse{
		Id: id,
	}, nil
}

func (k msgServer) InviteToGroup(goCtx context.Context, msg *types.MsgInviteToGroup) (*types.MsgInviteToGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	if !k.HasGroup(ctx, msg.Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Group))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Owner != k.GetGroupOwner(ctx, msg.Group) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Checks if the user is already in the group
	if k.AccountInGroup(ctx, msg.Invited, uint64(time.Now().UTC().Unix())) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Account already in group")
	}

	k.AddAccountGroup(ctx, msg.Invited, msg.Group, uint64(time.Now().UTC().Unix()))

	return &types.MsgInviteToGroupResponse{}, nil
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
