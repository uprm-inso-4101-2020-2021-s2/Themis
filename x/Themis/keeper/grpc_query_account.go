package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountAll(c context.Context, req *types.QueryAllAccountRequest) (*types.QueryAllAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accounts []*types.Account
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountStore := prefix.NewStore(store, types.KeyPrefix(types.AccountKey))

	pageRes, err := query.Paginate(accountStore, req.Pagination, func(key []byte, value []byte) error {
		var account types.Account
		if err := k.cdc.UnmarshalBinaryBare(value, &account); err != nil {
			return err
		}

		accounts = append(accounts, &account)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountResponse{Account: accounts, Pagination: pageRes}, nil
}

func (k Keeper) AccountWithNames(c context.Context, req *types.QueryAllAccountWithNamesRequest) (*types.QueryAllAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accounts []*types.Account
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountStore := prefix.NewStore(store, types.KeyPrefix(types.AccountNameKey))

	pageRes, err := types.PrefixPaginate(accountStore, types.KeyPrefix(req.Name), req.Pagination, func(key []byte, value []byte) error {
		var account types.Account
		if err := k.cdc.UnmarshalBinaryBare(value, &account); err != nil {
			return err
		}

		accounts = append(accounts, &account)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountResponse{Account: accounts, Pagination: pageRes}, nil
}

func (k Keeper) Account(c context.Context, req *types.QueryGetAccountRequest) (*types.QueryGetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var account types.Account
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAccount(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAccountIDBytes(req.Id)), &account)

	return &types.QueryGetAccountResponse{Account: &account}, nil
}

func (k Keeper) AccountAddress(c context.Context, req *types.QueryGetAccountAddressRequest) (*types.QueryGetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var account types.Account
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAccountAddr(ctx, req.Addr) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddrKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.GetStringBytes(req.Addr)), &account)

	return &types.QueryGetAccountResponse{Account: &account}, nil
}
