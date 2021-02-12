package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for Themis clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListPoll:
			return listPoll(ctx, k)
		case types.QueryListGroupPoll:
			return listGroupPolls(ctx, path[1:], k)
		case types.QueryGetPoll:
			return getPoll(ctx, path[1:], k)
		case types.QueryListAccount:
			return listAccount(ctx, k)
		case types.QueryGetAccount:
			return getAccount(ctx, path[1:], k)
		case types.QueryListUserAccount:
			return listUserAccounts(ctx, path[1:], k)
		case types.QueryListGroupAccount:
			return listGroupAccounts(ctx, path[1:], k)
		case types.QueryListGroup:
			return listGroup(ctx, k)
		case types.QueryGetGroup:
			return getGroup(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown Themis query endpoint")
		}
	}
}
