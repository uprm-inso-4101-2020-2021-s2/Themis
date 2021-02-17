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

func (k Keeper) VoteAll(c context.Context, req *types.QueryAllVoteRequest) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteKey))

	pageRes, err := query.Paginate(voteStore, req.Pagination, func(key []byte, value []byte) error {
		var vote types.Vote
		if err := k.cdc.UnmarshalBinaryBare(value, &vote); err != nil {
			return err
		}

		votes = append(votes, &vote)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoteResponse{Vote: votes, Pagination: pageRes}, nil
}

func (k Keeper) UserVoteAll(c context.Context, req *types.QueryAllUserVoteRequest) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountStore := prefix.NewStore(store, types.KeyPrefix(types.VotePtrKey))

	pageRes, err := PaginatePrefix(accountStore, types.KeyPrefix(types.VotePtrKey+req.User), req.Pagination, func(key []byte, value []byte) error {
		var vote types.Vote
		var votePtr types.VotePtr

		if err := k.cdc.UnmarshalBinaryBare(value, &votePtr); err != nil {
			return err
		}

		vote = k.GetVote(ctx, votePtr.Id)

		votes = append(votes, &vote)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoteResponse{Vote: votes, Pagination: pageRes}, nil
}

func (k Keeper) PollVoteAll(c context.Context, req *types.QueryAllPollVoteRequest) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountStore := prefix.NewStore(store, types.KeyPrefix(types.VoteKey))

	pageRes, err := PaginatePrefix(accountStore, types.KeyPrefix(types.VoteKey+req.Poll), req.Pagination, func(key []byte, value []byte) error {
		var vote types.Vote
		if err := k.cdc.UnmarshalBinaryBare(value, &vote); err != nil {
			return err
		}

		votes = append(votes, &vote)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoteResponse{Vote: votes, Pagination: pageRes}, nil
}

func (k Keeper) Vote(c context.Context, req *types.QueryGetVoteRequest) (*types.QueryGetVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vote types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.VoteKey+req.Id)), &vote)

	return &types.QueryGetVoteResponse{Vote: &vote}, nil
}
