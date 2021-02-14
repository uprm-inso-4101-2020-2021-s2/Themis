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

func (k Keeper) GroupAll(c context.Context, req *types.QueryAllGroupRequest) (*types.QueryAllGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var groups []*types.Group
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	groupStore := prefix.NewStore(store, types.KeyPrefix(types.GroupKey))

	pageRes, err := query.Paginate(groupStore, req.Pagination, func(key []byte, value []byte) error {
		var group types.Group
		if err := k.cdc.UnmarshalBinaryBare(value, &group); err != nil {
			return err
		}

		groups = append(groups, &group)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGroupResponse{Group: groups, Pagination: pageRes}, nil
}

func (k Keeper) Group(c context.Context, req *types.QueryGetGroupRequest) (*types.QueryGetGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var group types.Group
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.GroupKey+req.Id)), &group)

	return &types.QueryGetGroupResponse{Group: &group}, nil
}
