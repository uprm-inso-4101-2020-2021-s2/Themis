package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryGetPoll:
			return getPoll(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryListPoll:
			return listPoll(ctx, k, legacyQuerierCdc)

		case types.QueryGetAccount:
			return getAccount(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryListAccount:
			return listAccount(ctx, k, legacyQuerierCdc)

		case types.QueryListUserAccount:
			return listUserAccount(ctx, k, legacyQuerierCdc, path[1])

		case types.QueryListGroupAccount:
			return listGroupAccount(ctx, k, legacyQuerierCdc, path[1])

		case types.QueryGetGroup:
			return getGroup(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryListGroup:
			return listGroup(ctx, k, legacyQuerierCdc)

		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}
