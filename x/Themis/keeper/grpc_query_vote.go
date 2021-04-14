package keeper

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) VoteWithGroup(c context.Context, req *types.QueryVoteWithGroup) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteGroupKey))

	groupStr := strconv.FormatUint(req.Group, 10)
	pageRes, err := types.PrefixPaginate(voteStore, types.KeyPrefix(groupStr), req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) VoteWithGroupAndPoll(c context.Context, req *types.QueryVoteWithGroupAndPoll) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteGroupKey))

	groupStr := strconv.FormatUint(req.Group, 10)
	pollStr := strconv.FormatUint(req.Poll, 10)
	pageRes, err := types.PrefixPaginate(voteStore, types.KeyPrefix(groupStr+"-"+pollStr), req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) VoteWithUser(c context.Context, req *types.QueryVoteWithUser) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteUserKey))

	pageRes, err := types.PrefixPaginate(voteStore, types.KeyPrefix(req.User), req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) VoteWithUserAndPoll(c context.Context, req *types.QueryVoteWithUserAndPoll) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteUserKey))

	pollStr := strconv.FormatUint(req.Poll, 10)
	pageRes, err := types.PrefixPaginate(voteStore, types.KeyPrefix(req.User+"-"+pollStr), req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) VoteWithPoll(c context.Context, req *types.QueryVoteWithPoll) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteOptionKey))

	pollStr := strconv.FormatUint(req.Poll, 10)
	pageRes, err := types.PrefixPaginate(voteStore, types.KeyPrefix(pollStr), req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) VoteWithPollAndVote(c context.Context, req *types.QueryVoteWithPollAndVote) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var votes []*types.Vote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteStore := prefix.NewStore(store, types.KeyPrefix(types.VoteOptionKey))

	pollStr := strconv.FormatUint(req.Poll, 10)
	pageRes, err := types.PrefixPaginate(voteStore, types.KeyPrefix(pollStr+"-"+req.Vote), req.Pagination, func(key []byte, value []byte) error {
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

	if !k.HasVote(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetVoteIDBytes(req.Id)), &vote)

	return &types.QueryGetVoteResponse{Vote: &vote}, nil
}
