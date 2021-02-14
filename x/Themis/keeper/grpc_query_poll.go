package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PollAll(c context.Context, req *types.QueryAllPollRequest) (*types.QueryAllPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var polls []*types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKey))

	pageRes, err := query.Paginate(pollStore, req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.UnmarshalBinaryBare(value, &poll); err != nil {
			return err
		}

		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollResponse{Poll: polls, Pagination: pageRes}, nil
}

func (k Keeper) GroupPollAll(c context.Context, req *types.QueryAllGroupPollRequest) (*types.QueryAllPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var polls []*types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKey))

	pageRes, err := PaginatePrefix(pollStore, types.KeyPrefix(types.PollKey+req.Group), req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.UnmarshalBinaryBare(value, &poll); err != nil {
			return err
		}

		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollResponse{Poll: polls, Pagination: pageRes}, nil
}

func (k Keeper) Poll(c context.Context, req *types.QueryGetPollRequest) (*types.QueryGetPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var poll types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PollKey+req.Id)), &poll)

	return &types.QueryGetPollResponse{Poll: &poll}, nil
}
